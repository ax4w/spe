/*
* A virtual processor with 12 32bit registers for integer arithmetic
* R0-R7 are general use register, while R6 and R7 are used to store the top and lower 32bit of the multiplication
 */
package processor

import (
	"fmt"
	"spe/internal/memory"
	"strconv"
	"strings"
)

const (
	reg0 = iota
	reg1
	reg2
	reg3
	reg4
	reg5
	muld
	mulh
	pc
	sp
	ret
	param0
	param1
	ti
)

// word size is 32bit
type Processor struct {
	Registers [14]int32
	//FlRegisters          [6]float32
	translateIntRegister map[string]int32
	memory               *memory.Memory
}

func Init(m *memory.Memory) *Processor {
	p := &Processor{
		memory: m,
		translateIntRegister: map[string]int32{
			//for storing immediate when doing addi,..
			"ti": ti,
			//zero register
			"r0": reg0,
			//general usage
			"r1": reg1,
			"r2": reg2,
			"r3": reg3,
			"r4": reg4,
			"r5": reg5,
			"r6": muld,
			"r7": mulh,
			//specific
			"pc": pc,
			"sp": sp,
			"rt": ret,
			//function params
			"p0": param0,
			"p1": param1,
		},
	}
	p.Registers[9] = int32(m.Stack.Size)
	return p
}

func getArgs(s string, expected int) []string {
	args := strings.Split(s, ",")
	if len(args) != expected {
		panic(fmt.Sprintf("Expected arguments to be %d but got %d", expected, len(args)))
	}
	for i, v := range args {
		args[i] = strings.TrimSpace(v)
	}
	return args
}

func (p *Processor) toIntRegister(s string) int32 {
	if v, ok := p.translateIntRegister[strings.TrimSpace(strings.ToLower(s))]; !ok {
		panic(fmt.Sprintf("%s is not a valid register", s))
	} else {
		return v
	}
}

func (p *Processor) toInt(s string) int32 {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		if !strings.HasPrefix(s, "0x") {
			panic("value was not int or hex")
		}
		s = s[2:]
		val, err = strconv.ParseInt(s, 16, 32)
		if err != nil {
			panic(err.Error())
		}
	}
	return int32(val)

}

func (p *Processor) Reset() {
	for i := range p.Registers {
		if i == sp {
			p.Registers[i] = int32(p.memory.Stack.Size) - 1
		} else {
			p.Registers[i] = 0
		}
	}
}

func (p *Processor) incPC() {
	p.Registers[pc]++
}

func (p *Processor) Run() {
	println("running")
	println("got line loc: ", len(p.memory.Code))
	p.Registers[pc] = p.toInt(strings.Split(p.memory.Code[0], " ")[1])
	//p.Registers[pc] = p.memory.GetFunFromLabel("main")
	for p.Registers[pc] < int32(len(p.memory.Code)) {
		p.doInstruction(p.memory.CodeFromLine(p.Registers[pc]))
	}
}

func (p *Processor) getIntRegisterValue(reg int32) int32 {
	return p.Registers[reg]
}

/*
* Instruction syntax:
* -> double param: op param, param
* -> triple param: op param, param, param
* the op is 3 bytes long (3 chars)
 */
func (p *Processor) doInstruction(s string) {
	p.incPC()
	s = strings.ToLower(s)
	args := strings.Split(s, " ")
	op := args[0]
	if len(args) > 1 {
		s = s[len(op):]
	}
	switch op {
	case "lim":
		args := getArgs(s, 2)
		reg := p.toIntRegister(args[0])
		val := p.toInt(args[1])
		p.Registers[reg] = val
	case "mov":
		args := getArgs(s, 2)
		from := p.toIntRegister(args[0])
		to := p.toIntRegister(args[1])
		p.Registers[to] = p.getIntRegisterValue(from)
		p.Registers[from] = 0
	case "add":
		p.intOperation3(s, func(a, b int32) int32 { return a + b })
	case "sub":
		p.intOperation3(s, func(a, b int32) int32 { return a - b })
	case "mul":
		args := getArgs(s, 2)
		reg1 := p.toIntRegister(args[0])
		reg2 := p.toIntRegister(args[1])
		result := int64(p.Registers[reg1]) * int64(p.Registers[reg2])
		top := int32(result >> 32)
		bottom := int32(result & 0xFFFFFFFF)
		// Store results in designated registers
		p.Registers[mulh] = top
		p.Registers[muld] = bottom
	case "and":
		p.logical(s, func(a, b int32) int32 { return a & b })
	case "or":
		p.logical(s, func(a, b int32) int32 { return a | b })
	case "xor":
		p.logical(s, func(a, b int32) int32 { return a ^ b })
	case "bsl":
		p.logical(s, func(a, b int32) int32 { return a << b })
	case "bsr":
		p.logical(s, func(a, b int32) int32 { return a >> b })
	/*
	* function call
	 */
	case "call":
		args := getArgs(s, 1)
		line := p.toInt(args[0])
		p.Registers[pc] = line
	/*
	* jumping
	 */
	//unconditional jump
	case "jmp":
		p.jump(s)
	//if equal jump
	case "jeq":
		p.jumpWithFunc(s, func(a, b int32) bool { return a == b })
	//if a < b jump
	case "jlt":
		p.jumpWithFunc(s, func(a, b int32) bool { return a < b })
	//if a <= b jump
	case "jle":
		p.jumpWithFunc(s, func(a, b int32) bool { return a <= b })
	// if a > b jump
	case "jgt":
		p.jumpWithFunc(s, func(a, b int32) bool { return a > b })
	// if a >= b jump
	case "jge":
		p.jumpWithFunc(s, func(a, b int32) bool { return a >= b })
	//return
	case "ret":
		l := p.Registers[ret]
		println("returning to", l, "with", p.memory.Code[l])
		p.Registers[pc] = p.Registers[ret]
	/* stack interaction
	 */
	//store word stack
	case "sws":
		args := getArgs(s, 2)
		offset := p.toInt(args[0])
		reg := p.toIntRegister(args[1])
		addr := p.Registers[sp] + offset
		if addr >= int32(p.memory.Stack.Size) || addr < 0 {
			panic("out of bounds!")
		}
		p.memory.Stack.Set(addr, p.getIntRegisterValue(reg))
	case "lws":
		args := getArgs(s, 2)
		offset := p.toInt(args[0])
		reg := p.toIntRegister(args[1])
		addr := p.Registers[sp] + offset
		if addr >= int32(p.memory.Stack.Size) || addr < 0 {
			panic("out of bounds!")
		}
		p.Registers[reg] = p.memory.Stack.ReadDataAtAddr(addr)
	//zero out a int16 region on the stack
	case "nul":
		args := getArgs(s, 1)
		offset := p.toInt(args[0])
		addr := p.Registers[sp] + offset
		if addr < 0 || addr >= int32(p.memory.Stack.Size) {
			panic("address out of bounds")
		}
		p.memory.Stack.Set(addr, 0)
	//syscall
	case "exec":
		fn := p.Registers[param0]
		param := p.Registers[param1]
		p.exec(fn, param)
	case "nop":
	//debug
	case "!dump":
		println("register:")
		for i, v := range p.Registers {
			fmt.Printf("register[%d]: %d\n", i, v)
		}
	}
}
