package memory

import (
	"strings"
)

type stack struct {
	Data []int32
	Size uint16
}

type Memory struct {
	Code  []string //splitted at \n
	Stack stack
}

func New(size uint16) *Memory {
	data := make([]int32, size)
	return &Memory{
		Stack: stack{
			Data: data,
			Size: size,
		},
	}
}

func (m *Memory) LoadCode(s string) {
	lines := strings.Split(s, "\n")
	var processed []string
	//remove empty lines
	for _, v := range lines {
		if len(v) > 0 {
			processed = append(processed, v)
		}
	}
	lines = processed
	processed = processed[:0]
	//process code
	for _, v := range lines {
		processed = append(processed, strings.ToLower(v))
	}
	m.Code = processed
}

func (m *Memory) CodeFromLine(line int32) string {
	if line < 0 || line >= int32(len(m.Code)) {
		panic("Line is out of bouns")
	}
	return m.Code[line]
}

func (m *stack) ReadDataAtAddr(addr int32) int32 {
	if addr < 0 || addr >= int32(m.Size) {
		panic("Addr is not in the memory")
	}
	return m.Data[addr]
}

func (m *stack) Set(addr int32, val int32) {
	if addr < 0 || addr >= int32(m.Size) {
		panic("Addr is not in the memory")
	}
	m.Data[addr] = val
}
