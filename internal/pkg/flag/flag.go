package flag

import (
	"flag"
)

type flagVars struct {
	Fresh bool 
	Seeder bool 
}

var Flags *flagVars

func init() {
	fresh := flag.Bool("fresh", false, "Dropping all database tables before running new migration")
	seeder := flag.Bool("seeder", false, "Dropping all database tables before running new migration")

	flag.Parse()

	Flags = &flagVars{
		Fresh: *fresh,
		Seeder: *seeder,
	}
}	