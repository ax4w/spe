# Small Processor Emulator
`SPE` is a simple 12x 32Bit Register Stack-Only Processor Emulator to play with a very small and simple instruction set to get familiar with assembly and how to manage memory by maintaining the stack.

This processor emulator is desinged to do basic maths (`add`, `sub`, `mul`), bitwise functions (`and`, `or`, `xor`, `bsl`, `bsr`), branching (f.e `jump`, `jumpr`,...) and stack storage to write basic iterative or recursive calulcations (such as `fibonacci`). `Macros` get extended to the regular expressions to that **1 instruction = 1 cylce** (or per cycle one line of code)

## ToDo
[] add float support
[] add more functions
[] actually assembly it to binary

## Instructions
- `lim` (load immediate): `lim reg value`
- `mov`: `mov from to`
    - moves from to to and sets from to zero
- `add`: `add store reg1 reg2`
- `sub`: `sub store reg1 reg2`
- `mul`: `mul reg1 reg2`
    - top 32bit is in R7 and bottom 32bit in R6
- `and`: `and store reg1 reg2`
- `xor`: `xor store reg1 reg2`
- `or`: `or store reg1 reg2`
- `bsl`: `bsl store reg1 reg2`
    - bitshift reg1::value by reg2::value bits to the left
- `bsr`: `bsr store reg1 reg2 `
    - bitshift reg1::value by reg2::value bits to the right
- `jmp`: `jmp label`
- `jeq` (jump if equal): `jeq reg1 reg2 label`
    - reg1 == reg2
- `jlt` (jump less than): `jlt reg1 reg2 label`
    - reg1 < reg2
- `jle` (jump if less than or equal): `jle reg1 reg2 label`
    - reg1 <= reg2
- `jgt` (jump if greater than): `jgt reg1 reg2 label`
    - reg1 > reg2
- `jge` (jump if greater than or equal): `jge reg1 reg2 label`
    - reg1 >= reg2
- `ret` (return): `ret`
    - returns to the address stored in `rt`
- `sws` (store word stack): `sws offset_from_sp reg`
- `lws` (load word stack): `lws offset_from_sp reg`
- `exec` (execute)
    - calls function id (param0) with argument (param1)

## Register
### general usage
- `r0` - `r7`
> Important: r0 is the zero register and should not be set
### specific
- `sp` (stack pointer)
- `pc` (programm counter)
- `rt` (return address)
### paramters
- `param0` - `param1`
### Internal Registers
While `r6` and `r7` are used for the result of `mul` its still ok to use it to store data.
`SPE` has an `ti` register to store the immediate of the `I-Macros` for the operation. This should not be used to store data.

## Global Values
Global values can be defined using `.global <name_no_spaces> <value_int_or_hex>`. Globals are be defined before any instructions, because of the code gen.

Global values can be used in instructions using `.<name_no_spaces>` and are handled by the code gen with subsitution

## Macros
### R-suffix
> All instructions ending with an 'r' set the ret register to current-address +1
- `jumpr` (jump and register): `jumpr label`
- `jeqr` (jump equals and register): `jeqr reg reg label`
- `jltr` (jump less than and register): `jltr reg reg label`
- `jler` (jump less than or equal and register): `jler reg reg label`
- `jgtr` (jump greater than and register): `jgtr reg reg label`
- `jger` (jump greater than or equal and register): `jger reg reg label`
### Stack
- `push` (pushes register to the stack): `push reg`
    - automatically increments stack by 1x int32
- `pop` (pops the top value to the register): `pop reg`
    - automatically decrements stack by 1x int32
### I-Suffix
> All instructions ending with an 'i' take an immediate as one paramter
- `addi` (add immediate): `addi store reg value`
- `subi` (add immediate): `subi store reg value`
- `jei`  (jump equal immediate): `jei reg value label `
- `bsli`: `bsl store reg1 value`
    - bitshift reg1::value by value bits to the left
- `bsri`: `bsr store reg1 value`
    - bitshift reg1::value by values bits to the right
## Functions
- `print`: ID 0 and 1 Parameter

## Labels and Functions
### Labels `LABEL`
With the jump-type instructions you can jump to labels. Labels should be used for branching

### Functions `FUN`
With the call instruction you can call a function. Call sets the `rt`-register and jumps to the function.
Functions should be used to delclare contained blocks of logic.
