package main

import (
	"flag"
	"fmt"
	"glitter/commands"
	"glitter/tools"
	"os"
)

var commitCmd = flag.NewFlagSet("commit", flag.ExitOnError)
var commitMsg = commitCmd.String("m", "", "Commit Message")

func main() {
	tools.OpSys()
	if len(os.Args) < 2 {
		fmt.Println("Expected a command!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		commands.Init()
	case "status":
		commands.Status()
	case "add":
		commands.Add(os.Args[2])
	case "commit":
		commitCmd.Parse(os.Args[2:])
		commands.Commit(*commitMsg)
	}
}
