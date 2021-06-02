package main

import (
  "strings"
  "bufio"
  "os"
  "log"
)


type Parser struct {
  scanner *bufio.Scanner
  source string
  current_line string
  completed bool
  counter int
}


func (p *Parser) HasMoreCommands() bool {
  return !p.completed
}


func (p *Parser) Advance() bool {
  if p.scanner == nil {
    file, err := os.Open(p.source)

    if err != nil {
      log.Fatal(err)
    }

    p.scanner = bufio.NewScanner(file)
    p.counter = 0
  }

  if err := p.scanner.Err(); err != nil {
    log.Fatal(err)
  }

  p.current_line = ""

  if p.scanner.Scan() {
    p.counter++
    line := p.scanner.Text()

    if len(line) > 0 {
      trimmed := strings.TrimSpace(line)

      if !strings.HasPrefix(trimmed, "//") && trimmed != "" {
        p.current_line = trimmed
      }
    }

  } else {
    p.completed = true
    return false
  }

  return true
}


func (p *Parser) CommandType() string {
  if strings.HasPrefix(p.current_line, "(") {
    return "L_COMMAND"

  } else if strings.HasPrefix(p.current_line, "@") {
    return "A_COMMAND"

  } else if p.current_line != "" {
    return "C_COMMAND"
  }

  return ""
}


func (p *Parser) Symbol() string {
  switch command := p.CommandType(); command {
    case "L_COMMAND":
      return strings.Split(strings.Split(p.current_line, "(")[1], ")")[0]
    case "A_COMMAND":
      return strings.Split(p.current_line, "@")[1]
  }

  return ""
}


func (p *Parser) Dest() string {
  if command := p.CommandType(); command != "C_COMMAND" {
    return ""
  }

  arr := strings.Split(p.current_line, "=")

  if len(arr) > 1 {
    return arr[0]
  }

  return ""
}


func (p *Parser) Comp() string {
  if command := p.CommandType(); command != "C_COMMAND" {
    return ""
  }

  arr := strings.Split(p.current_line, "=")
  out := strings.Split(arr[0], ";")[0]

  if len(arr) > 1 {
    out = strings.Split(arr[1], ";")[0]
  }

  return strings.TrimSpace(strings.Split(out, "//")[0])
}


func (p *Parser) Jump() string {
  if command := p.CommandType(); command != "C_COMMAND" {
    return ""
  }

  arr := strings.Split(p.current_line, ";")

  if len(arr) > 1 {
    return strings.TrimSpace(strings.Split(arr[1], "//")[0])
  }

  return ""
}


func createParser(source string) Parser {
  return Parser{source: source, current_line: "", completed: false}
}
