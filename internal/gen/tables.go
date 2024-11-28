package gen

func tables(lines []string) ([]string, map[string]int, map[string]int) {
	var (
		linesLabel = make(map[string]int)
		linesFun   = make(map[string]int)
		result     []string
	)
	for i, v := range lines {
		inst, op, args := extractOpAndArgs(v)
		switch op {
		case "label":
			expect(inst, 1, len(args))
			linesLabel[args[0]] = i + 1
			result = append(result, "nop")
		case "fun":
			expect(inst, 1, len(args))
			linesFun[args[0]] = i + 1
			result = append(result, "nop")
		default:
			result = append(result, inst)
		}
	}
	return result, linesLabel, linesFun
}
