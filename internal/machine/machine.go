package machine

import (
	"errors"
	"fmt"
	"os"
	"spe/internal/memory"
	"spe/internal/processor"
)

type Machine struct {
	processor *processor.Processor
	memory    *memory.Memory
}

func New(stackSize uint16) *Machine {
	mem := memory.New(stackSize)
	return &Machine{
		processor: processor.Init(mem),
		memory:    mem,
	}
}

func (m *Machine) Processor() *processor.Processor {
	return m.processor
}

func (m *Machine) loadFromFile(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%s is not a valid path to load from", path)
		return
	}
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("An error occured while loading the file: %v", err)
	}
	dataAsString := string(data)
	m.memory.LoadCode(dataAsString)
	m.processor.Reset() //reset, because this function can be called after code was executed
}

func (m *Machine) Run(path string) {
	m.loadFromFile(path)
	m.processor.Run()
}
