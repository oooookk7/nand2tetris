// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// -------------------------------------
// Pseudocode:
//
// counter = R0
//
// while (counter != 0):
//   R2 += R1
//   counter -= 1
//
// Steps:
//   1. Set R2 to 0 and the counter to R0.
//   2. Load the counter to D and check if it equals 0. If so, jump to to end of loop.
//   4. Load R1 into D.
//   5. Increment R2 with D.
//   6. Decrement the counter with 1.
//   7. Re-run steps 2 to 6.
//   8. Set the code to end (infinitely). Infinite loop is the standard way to terminate execution of Hack programs to prevent running into malicious code.
//
// Shortcuts:
//   A: Address register. Holds the current address and can be manipulated as it acts as a pointer to the specified address.
//   M: Main memory. Holds the RAM address contents for the @VAR.
//   D: Data register. Holds the operation read from memory location.
// -------------------------------------

@R2
M=0

@R0
D=M

@counter
M=D

(LOOP)
  @counter
  D=M

  @END
  D;JEQ

  @R1
  D=M

  @R2
  M=M+D

  @counter
  M=M-1

  @LOOP
  0;JMP

(END)
  @END
  0;JMP
