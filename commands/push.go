package commands

import (
	"fmt"
	"os/exec"
)

func Push() {
	cmd := exec.Command("git", "push")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
	fmt.Println("Pushed!")
}
