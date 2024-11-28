package gen

import "fmt"

func fixInstructionWithLabels(instructions []string) []string {
	var result []string
	l, lines, funs := tables(instructions)
	for _, v := range l {
		inst, op, args := extractOpAndArgs(v)
		switch op {
		case "call":
			expect(op, 1, len(args))
			if fl, ok := funs[args[0]]; !ok {
				panic(fmt.Sprintf("no fund named %s found", args[0]))
			} else {
				result = append(result, fmt.Sprintf("%s %d", op, fl))
			}
		case "jmp":
			expect(op, 1, len(args))
			if fl, ok := lines[args[0]]; !ok {
				panic("no fun found")
			} else {
				result = append(result, fmt.Sprintf("%s %d", op, fl))
			}
		case "jeq", "jlt", "jle", "jgt", "jge":
			expect(op, 3, len(args))
			if fl, ok := lines[args[2]]; !ok {
				panic("no fun found")
			} else {
				result = append(result, fmt.Sprintf("%s %s, %s, %d", op, args[0], args[1], fl))
			}
		default:
			result = append(result, inst)
		}
	}
	if v, ok := funs["main"]; !ok {
		panic("no main function definded")
	} else {
		result = append([]string{fmt.Sprintf("entry %d", v)}, result...)
	}
	return result
}
