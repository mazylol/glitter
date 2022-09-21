package main

import (
	"flag"
	"fmt"
	"glitter/commands"
	"glitter/tools"
	"os"
)

func main() {
	tools.OpSys()

	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)
	commitMsg := commitCmd.String("m", "", "Commit Message")

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
		fmt.Println("  message: ", *commitMsg)
	}
}
