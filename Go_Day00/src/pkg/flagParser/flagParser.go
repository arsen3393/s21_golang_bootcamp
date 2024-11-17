package flagParser

import (
	"flag"
)

type Flags struct{
	Mean bool
	Median bool
	Mode bool
	SD bool
}


func Parse() (UserFlags Flags){
	flag.BoolVar(&UserFlags.Mean, "Mean", false, "Mean flag")
	flag.BoolVar(&UserFlags.Median, "Median", false, "Median flag")
	flag.BoolVar(&UserFlags.Mode, "Mode", false, "Mode flag")
	flag.BoolVar(&UserFlags.SD, "SD", false, "Flags flag")
	flag.Parse()
	CheckFlagsOnFalse(&UserFlags)
	return
}

func CheckFlagsOnFalse(UserFlags *Flags){
	if !(UserFlags.Mean || UserFlags.Median || UserFlags.Mode || UserFlags.SD){
		UserFlags.Mean = true
		UserFlags.Median = true
		UserFlags.Mode = true
		UserFlags.SD = true
	}
}

