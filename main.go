package main

import (
	"fmt"
	"glitter/commands"
	"glitter/tools"
	"os"
)

func main() {
	tools.OpSys()
	if len(os.Args) < 2 {
		fmt.Println("Expected a command!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	}
}
