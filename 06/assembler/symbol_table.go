package main

import (
  "fmt"
)


type SymbolTable struct {
  symbols map[string]int
  address_counter int
}


func (st *SymbolTable) AddEntry(symbol string, address int) {
  st.symbols[symbol] = address
}


func (st *SymbolTable) Contains(symbol string) bool {
  if _, exists:= st.symbols[symbol]; exists {
    return true
  }

  return false
}


func (st *SymbolTable) GetAddress(symbol string) int {
  if st.Contains(symbol) {
    return st.symbols[symbol]
  }

  return -1
}


func (st *SymbolTable) NextAddress() int {
  // Use within ranges of 16 - 16383, 16385 - 24575 and 24577 - 32767
  next_counter := st.address_counter + 1

  if next_counter == 16384 || next_counter == 24576 {
    st.address_counter = next_counter + 1

  } else {
    st.address_counter = next_counter
  }

  return st.address_counter
}


func createSymbolTable() SymbolTable {
  symbols := make(map[string]int)

  for i := 0; i < 16; i++ {
    symbols[fmt.Sprint("R", i)] = i
  }

  symbols["SP"] = 0
  symbols["LCL"] = 1
  symbols["ARG"] = 2
  symbols["THIS"] = 3
  symbols["THAT"] = 4
  symbols["SCREEN"] = 16384
  symbols["KBD"] = 24576

  return SymbolTable{symbols: symbols, address_counter: 15}
}

