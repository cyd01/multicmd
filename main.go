package main

import (
	"fmt"
	"os"
	
	"multicmd/cmd1"
	"multicmd/cmd2"
)

func usage() {
	fmt.Println("Available commands: cmd1, cmd2")
}

func main() {
	if len(os.Args)<= 1 {
		usage()
	} else {
		argsWithoutProg := os.Args[1:]
		switch cmd:=os.Args[1]; cmd {
			case "cmd1":
				cmd1.Main(argsWithoutProg)
			case "cmd2":
				cmd2.Main(argsWithoutProg)
			default:
				fmt.Println("Unknown command "+cmd)
				usage()
		}
	}
}
