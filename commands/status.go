package commands

import (
	"fmt"
	"os/exec"
)

func Status() {
	cmd := exec.Command("git", "status")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
