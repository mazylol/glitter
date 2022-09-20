package tools

import (
	"fmt"
	"os"
	"runtime"
)

func OpSys() {
	rt := runtime.GOOS
	switch rt {
	case "windows":
		fmt.Println("Windows is not supported!")
		os.Exit(1)
	case "darwin":
		fmt.Println("Darwin is not supported!")
		os.Exit(1)
	default:

	}
}
