package gen

/*
* Globals are not needed when executing, since they can be subsituted like #define in c
* we get all globals, which are (convention) defined before any instruction
 */

import (
	"fmt"
	"strings"
)

func replaceGlobals(lines []string) []string {
	var (
		result        []string
		globalToValue = make(map[string]string)
		i             = 0
	)
	for ; i < len(lines) && strings.HasPrefix(lines[i], ".global"); i++ {
		line := lines[i]
		args := strings.Split(line, " ")
		expect(line, 3, len(args))
		globalToValue["."+args[1]] = args[2]
	}
	lines = lines[i:]
	for _, v := range lines {
		//only immediate operations can reference globals
		inst, op, args := extractOpAndArgs(v)
		switch op {
		case "lim":
			//<op> <reg> <value>
			expect(inst, 2, len(args))
			value := args[1]
			if strings.HasPrefix(value, ".") {
				if v, ok := globalToValue[value]; ok {
					result = append(result, fmt.Sprintf("%s %s, %s", op, args[0], v))
				} else {
					panic(fmt.Sprintf("global value %s was not found", value))
				}
			} else {
				result = append(result, inst)
			}
		case "addi", "subi", "bsri", "bsli":
			//<op> <reg>, <reg>, value
			expect(inst, 3, len(args))
			value := args[2]
			if strings.HasPrefix(value, ".") {
				if v, ok := globalToValue[value]; ok {
					result = append(result, fmt.Sprintf("%s %s, %s, %s", op, args[0], args[1], v))
				} else {
					panic(fmt.Sprintf("global value %s was not found", value))
				}
			} else {
				result = append(result, inst)
			}
		case "jei", "jeir":
			//<op> <reg>, <value>, <label>
			expect(inst, 3, len(args))
			value := args[1]
			if strings.HasPrefix(value, ".") {
				if v, ok := globalToValue[value]; ok {
					result = append(result, fmt.Sprintf("%s %s, %s, %s", op, args[0], v, args[2]))
				} else {
					panic(fmt.Sprintf("global value %s was not found", value))
				}
			} else {
				result = append(result, inst)
			}
		default:
			result = append(result, inst)
		}
	}
	return result
}
