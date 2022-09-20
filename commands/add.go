package commands

import (
	"fmt"
	"os/exec"
)

func Add(arg0 string) {
	cmd := exec.Command("git", "add", arg0)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
