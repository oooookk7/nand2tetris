// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed.
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// -------------------------------------
// Pseudocode:
//
// while (1):
//   color = 0;
//
//   if (KBD > 0):
//     color = 1;
//
//   index = SCREEN;
//   counter = 24576 - SCREEN;
//
//   while (counter != 0):
//     RAM[index] = color;
//     counter -= 1;
//     index += 1;
//
// White: 0, Black: 1
// -------------------------------------

@color
M=0

@index
M=0

@counter
M=0

(DETECT)
  @24576
  D=A

  @counter
  M=D

  @SCREEN
  D=A

  @index
  M=D

  @counter
  M=M-D

  @KBD
  D=M

  @PRESSED
  D;JGT

  @color
  M=0

  @DRAW
  0;JMP

(PRESSED)
  @color
  M=1

(DRAW)
  @counter
  D=M

  @DETECT
  D;JMP

  @index
  D=M

  @pixel
  M=D

  @color
  D=M

  @pixel
  A=M
  M=D

  @index
  M=M+1

  @counter
  M=M-1
  D=M

  @DRAW
  0;JMP