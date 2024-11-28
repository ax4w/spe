package gen

/*
* This file contains all the methods to expand the macros to the regular instructions
 */

import (
	"fmt"
	"strings"
)

func expect(inst string, a, b int) {
	if a != b {
		panic(fmt.Sprintf("%s expects %d but got %d", inst, a, b))
	}
}

func extractOpAndArgs(inst string) (string, string, []string) {
	args := strings.Split(inst, " ")
	op := args[0]
	cutInst := inst
	if len(args) > 1 {
		cutInst = inst[len(op):]
	}
	args = strings.Split(cutInst, ",")
	for i, v := range args {
		args[i] = strings.TrimSpace(strings.ToLower(v))
	}
	return inst, op, args
}

func expand(inst string) []string {
	inst, op, args := extractOpAndArgs(inst)
	switch op {
	case "addi":
		//addi <reg>, <reg>, val
		expect(op, 3, len(args))
		return []string{
			fmt.Sprintf("lim ti, %s", args[2]),
			fmt.Sprintf("add %s, %s, ti", args[0], args[1]),
		}
	case "subi":
		//addi <reg>, <reg>, val
		expect(op, 3, len(args))
		return []string{
			fmt.Sprintf("lim ti, %s", args[2]),
			fmt.Sprintf("sub %s, %s, ti", args[0], args[1]),
		}
	case "call":
		//call <name>
		expect(op, 1, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("call %s", args[0]),
		}
	case "jei":
		//jei <reg>, <value>, <label>
		expect(op, 3, len(args))
		return []string{
			fmt.Sprintf("lim ti, %s", args[1]),
			fmt.Sprintf("jeq %s, ti, %s", args[0], args[2]),
		}
	case "jeir":
		//jeir <reg>, <value>, <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("lim ti, %s", args[1]),
			fmt.Sprintf("jeq %s, ti, %s", args[0], args[2]),
		}
	case "jmpr":
		//jeir <label>
		expect(op, 1, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jmp %s", args[0]),
		}
	case "jeqr":
		//jeir <reg> <reg> <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jeq %s, %s, %s", args[0], args[1], args[2]),
		}
	case "jltr":
		//jltr <reg> <reg> <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jlt %s, %s, %s", args[0], args[1], args[2]),
		}
	case "jler":
		//jler <reg> <reg> <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jle %s, %s, %s", args[0], args[1], args[2]),
		}
	case "jgtr":
		//jgtr <reg> <reg> <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jgt %s, %s, %s", args[0], args[1], args[2]),
		}
	case "jger":
		//jger <reg> <reg> <label>
		expect(op, 3, len(args))
		return []string{
			"lim ti, 1",
			"add rt, pc, ti",
			fmt.Sprintf("jge %s, %s, %s", args[0], args[1], args[2]),
		}
	case "push":
		//push <reg>
		expect(op, 1, len(args))
		return []string{
			"lim ti, 1",
			"sub sp, sp, ti",
			fmt.Sprintf("sws 0, %s", args[0]),
		}
	case "pop":
		//pop <reg>
		expect(op, 1, len(args))
		return []string{
			fmt.Sprintf("lws 0, %s", args[0]),
			"lim ti, 1",
			"add sp, sp, ti",
		}
	case `bsri`:
		//bsri <reg>, <reg>, val
		expect(op, 3, len(args))
		return []string{
			fmt.Sprintf("lim ti, %s", args[2]),
			fmt.Sprintf("bsr %s, %s, ti", args[0], args[1]),
		}
	case `bsli`:
		//bsri <reg>, <reg>, val
		expect(op, 3, len(args))
		return []string{
			fmt.Sprintf("lim ti, %s", args[2]),
			fmt.Sprintf("bsl %s, %s, ti", args[0], args[1]),
		}
	default:
		return []string{inst}
	}
}
