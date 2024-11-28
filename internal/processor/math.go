package processor

func (p *Processor) intOperation3(inst string, fn func(a, b int32) int32) {
	args := getArgs(inst, 3)
	regToStoreTo := p.toIntRegister(args[0])
	reg1 := p.toIntRegister(args[1])
	reg2 := p.toIntRegister(args[2])
	p.Registers[regToStoreTo] = fn(p.getIntRegisterValue(reg1), p.getIntRegisterValue(reg2))
}
