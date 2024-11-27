package flag

import (
	"flag"
)

type flagVars struct {
	Fresh bool 
}

var Flags *flagVars

func init() {
	fresh := flag.Bool("fresh", false, "Dropping all database tables before running new migration")

	flag.Parse()

	Flags = &flagVars{
		Fresh: *fresh,
	}
}	