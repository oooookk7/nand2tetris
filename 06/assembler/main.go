package main

import (
  "fmt"
  "unicode"
  "strconv"
  "os"
  "log"
  "strings"
)


type Assembler struct {
  symbol_table SymbolTable
  contents []string
  source string
  target string
}


func (a *Assembler) FirstPass () {
  p := createParser(a.source)
  line_number := 0

  for has_line := true; has_line && !p.completed; has_line = p.Advance() {
    command_type := p.CommandType()

    if p.current_line != "" {
      if command_type == "L_COMMAND" {
        a.symbol_table.AddEntry(p.Symbol(), line_number)

      } else {
        line_number++
      }
    }
  }
}


func (a *Assembler) SecondPass () {
  p := createParser(a.source)

  for p.Advance() {
    if p.current_line != "" {
      command_type := p.CommandType()

      if (command_type == "L_COMMAND") {
        continue
      }

      if (command_type == "A_COMMAND") {
        symbol := p.Symbol()
        address := a.symbol_table.GetAddress(symbol)

        if unicode.IsDigit(rune(symbol[0])) {
          address, _ = strconv.Atoi(symbol)

        } else if address == -1 {
          address = a.symbol_table.NextAddress()
          a.symbol_table.AddEntry(symbol, address)
        }

        a.contents = append(a.contents, "0" + fmt.Sprintf("%015s", strconv.FormatInt(int64(address), 2)))

      } else {
        c := createCode()
        a.contents = append(a.contents, "111" + fmt.Sprintf("%013s", c.Comp(p.Comp()) + c.Dest(p.Dest()) + c.Jump(p.Jump())))
      }
    }
  }
}


func (a *Assembler) Compile () {
  file, err := os.Create(a.target)

  if err != nil {
    log.Fatal(err)
  }

  file.WriteString(strings.Join(a.contents, "\n"))
  file.Close()
}


func main() {
  source := os.Args[1]
  target := strings.Split(source, ".asm")[0] + ".hack"
  a := Assembler{symbol_table: createSymbolTable(), source: source, target: target, contents: []string {}}

  fmt.Print("Compiling to ", target, "\n")
  a.FirstPass()
  a.SecondPass()
  a.Compile()
}
