package compile

import (
	"fmt"
	"os/exec"
	"runtime"
)

func GetOs() (string, bool) {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windowsss")
		cmd := exec.Command("gcc", "--version")
		out, err := cmd.Output()
		if err == nil {
			fmt.Println(string(out))
			return string(out), true
		}

	}
	return "", false
}
