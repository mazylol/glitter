package commands

import (
	"fmt"
	"os/exec"
)

func Commit(msg string) {
	cmd := exec.Command("git", "commit", "-m", msg)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}
