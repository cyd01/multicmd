package cmd1

import (
	"fmt"
	"os"
	"strconv"
	
	"multicmd/flags"
)

type Empty struct{}

const (
	command = "cmd1"
)

var (
	f = flags.NewFlag( command )
	debug = f.Bool( "debug", false, "Set debug mode" )
	val1 = f.String( "st", "test", "Set stage" )
	val2 = f.Int( "pos", 421, "Set position" )
)

func usage() {
	fmt.Println("This is the usage of "+command)
	os.Exit(0)
}

func Main(args []string) {
	fmt.Println("in "+command)
	f.SetUsage(usage)
	f.Parse(args[1:])
	
	fmt.Println("Here is the value: "+*val1)
	fmt.Println("Here is the position: "+strconv.Itoa(*val2))
}
