package processor

import "fmt"

func (p *Processor) exec(fn int32, param any) {
	switch fn {
	//print
	case 1:
		fmt.Printf("%v\n", param)
	}
}
