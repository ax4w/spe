package gen

import (
	"os"
	"strings"
)

func File(path string) {
	var (
		instructions []string
	)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	lines := strings.Split(string(bytes), "\n")
	for _, v := range lines {
		v = strings.ToLower(v)
		if len(v) > 0 {
			instructions = append(instructions, expand(v)...)
		} else {
			//replace empty lines with nop
			instructions = append(instructions, "nop")
		}
	}
	instructions = replaceGlobals(instructions)
	instructions = fixInstructionWithLabels(instructions)
	os.WriteFile("output.tspe", []byte(strings.Join(instructions, "\n")), os.ModePerm)
}
