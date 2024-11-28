package processor

func (p *Processor) jump(inst string) {
	args := getArgs(inst, 1)
	line := p.toInt(args[0])
	p.Registers[pc] = line
}

func (p *Processor) jumpWithFunc(inst string, fn func(a, b int32) bool) {
	args := getArgs(inst, 3)
	reg1 := p.getIntRegisterValue(p.toIntRegister(args[0]))
	reg2 := p.getIntRegisterValue(p.toIntRegister(args[1]))
	line := p.toInt(args[2])
	if fn(reg1, reg2) {
		p.Registers[pc] = line
	}
}
