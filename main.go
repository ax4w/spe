package main

import (
	"flag"
	"spe/internal/gen"
	"spe/internal/machine"
)

var (
	ftranslate       = flag.Bool("t", false, "tranlsate file to output")
	frun             = flag.Bool("r", false, "tranlsate file to output")
	ftranslateAndRun = flag.Bool("tr", false, "tranlsate file to output")
	ffile            = flag.String("p", "", "path to file")
	fstackSize       = flag.Int("s", 1024, "stack size (1 = 1x int32)")
)

func init() {
	flag.Parse()
	if !*ftranslate && !*frun && !*ftranslateAndRun {
		panic("nothing to do!")
	}
	if *fstackSize == 0 {
		panic("stackSize cannot be 0")
	}
	if len(*ffile) == 0 {
		panic("no file provided")
	}
}

func main() {
	if *ftranslate {
		gen.File(*ffile)
	} else if *frun {
		m := machine.New(uint16(*fstackSize))
		m.Run(*ffile)
	} else if *ftranslateAndRun {
		gen.File(*ffile)
		m := machine.New(uint16(*fstackSize))
		m.Run("./output.tspe")
	}
}
