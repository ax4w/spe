package processor

// every logical op, that uses this func, has 3 args
func (p *Processor) logical(inst string, fn func(a, b int32) int32) {
	args := getArgs(inst, 3)
	target := p.toIntRegister(args[0])
	reg1 := p.toIntRegister(args[1])
	reg2 := p.toIntRegister(args[2])
	p.Registers[target] = fn(p.getIntRegisterValue(reg1), p.getIntRegisterValue(reg2))
}
