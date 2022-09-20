package commands

import (
	"fmt"
	"os/exec"
)

func Init() {
	cmd := exec.Command("git", "init", ".")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
