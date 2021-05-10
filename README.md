# Nand2Tetris

This project contains project codes learnt from the self-paced course [nand2tetris](https://www.nand2tetris.org/course).

> Ensure that you have download the related project software tools [here](https://www.nand2tetris.org/software).

### [Project 1](https://www.nand2tetris.org/project01): Boolean Logic

Focuses on the basic boolean logic gates. For example, XOR, OR and AND gates, which are the fundamental building blocks. 

**How to navigate the project:**

1. The boolean table for each logic chip can be found in the related `.cmp` files.
2. The hardware description language (HDL) can be found in the related `.hdl` files.
3. The test scripts are found in the `.tst` files. To run them, run the `./HardwareSimulator.sh` file, click on "Load Script" and select the related `.tst` file.

The primitive chips can be found in the tutorial for "Hardware Simulator" in [here](https://www.nand2tetris.org/software).


### [Project 2](https://www.nand2tetris.org/project02): Boolean Arithmetic

Focuses on the arithmetic logic gates. For example, a basic ALU (Arithmetic Logic Unit) that does basic arithmetic and logic operations in the CPU, which contains as well a CU (Control Unit) that directs the operation of the processor for the ALUs (e.g. what to compute).

_See Project 1 on how to navigate the project._


### [Project 3](https://www.nand2tetris.org/project03): Sequential Logic

Focuses on the sequential logic gates. For example, a 1-bit register state is sequential (e.g. at time = 2 the current stored state is the output of time = 1), and it contains a load flag that enables writes to the bit register, else the state is locked. You could imagine an ever-looping circuit that determines if state changes at every discrete time frame (e.g. `out(t) = in(t-1)`. These bits can be found in the CPU processor register, or RAM memory.

_See Project 1 on how to navigate the project._


### [Project 4](https://www.nand2tetris.org/project04): Machine Language

Focuses on the assembly language, which is compiled to machine language. The languages are compiled in such order: high-level > assembly-level > machine-level.

For example, in a high-level code:

```
int c = a * b;
```

Then in an assembly-level code:

```
MOV CX, [AX] ; Moves AX to CX
MUL BX;      ; Result in CX = AX * BX
```

Then in a machine-level compiled code:

```
0000000000000010
1110101010001000
0000000000000001
...
```

The CPU, RAM and Input/Output devices (e.g. keyboard, mouse, hard-disk, computer screen) are connected to a system bus that carries information, determine where a data instruction should be sent. The CPU utilitises this bus to interact with the RAM and Input/Output to do operations.

Another way to look at it is an assembly-level command `ADD R1, R2, R9`, which represents an addition operation from R1 and R2 to be stored into R9 gets compiled into a machine-level code `1010 001100 011001` (where `1010` represents `ADD`, `001100` R1, `0001` R2 and `011001` R3).

It is vital to understand the registers [within the CPU](https://computersciencewiki.org/index.php/Registers_within_the_CPU).

**How to navigate the project:**

1. The assembly scripts are tagged as `.asm` files.
2. The result table for each assembly script is stored in the relevant  `.cmp` files.
3. The test scripts are found in the `.tst` files. To run them,
    1. Run `./Assembler.sh` file, click on "Load Source file" and select the related `.asm` file. Run the script and save the file the output as the relevant `.hack` file.
    2. Run `./CPUEmulator.sh` file, click on "Load Script" and select the related `.tst` file.
    3. If there are any errors, restart the steps from 1 onward.

The Hack assembly commands can be found in the PDF tutorial for "Project 4" in [here](https://www.nand2tetris.org/course).


### [Project 5](https://www.nand2tetris.org/project05): Computer Architecture

Focuses on the how the RAM, screen, keyboard, memory and CPU interacts with each other as a computer. For example, in the CPU it contains the Program Counter (PC), ALU, Data Register (DR) and Address Register (AD). Each instruction given to the CPU is either stored or fetched to/from the AD or DR, and then passed to the ALU to execute, after which (if it's an AD fetch) the PC increments (sometimes it may skip if command is jump), executing the next instruction.

_See Project 1 on how to navigate the project._

