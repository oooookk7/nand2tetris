# Nand2Tetris

This project contains project codes learnt from the self-paced course [nand2tetris](https://www.nand2tetris.org/course).

### Project 1: Boolean Logic

Focuses on the basic boolean logic gates. For example, XOR, OR and AND gates, which are the fundamental building blocks. 

**How to navigate the project:**

1. The boolean table for each logic gate can be found in the related `.cmp` files.
2. The hardware description language (HDL) can be found in the related `.hdl` files.
3. The test scripts are found in the `.tst` files. To run them, download the related project software tools and run the `./HardwareSimulator.sh` file, click on "Load Script" and select the related `.tst` file.


### Project 2: Boolean Arithmetic

Focuses on the arithmetic logic gates. For example, a basic ALU (Arithmetic Logic Unit) that does basic arithmetic and logic operations in the CPU, which contains as well a CU (Control Unit) that directs the operation of the processor for the ALUs (e.g. what to compute).

_See Project 1 on how to navigate the project._


### Project 3: Sequential Logic

Focuses on the sequential logic gates. For example, a 1-bit register state is sequential (e.g. at time = 2 the current stored state is the output of time = 1), and it contains a load flag that enables writes to the bit register, else the state is locked. You could imagine an ever-looping circuit that determines if state changes at every discrete time frame (e.g. `out(t) = in(t-1)`. These bits can be found in the CPU processor register, or RAM memory.

_See Project 1 on how to navigate the project._









