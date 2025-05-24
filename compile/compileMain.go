package compile

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var outputFileAddress string = "output"

func GetOs() (string, bool) {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windowsss")
		cmd := exec.Command("gcc", "--version")
		out, err := cmd.Output()
		if err == nil {
			fmt.Println(string(out))
			compilerToExe(nil)
			return string(out), true
		}
		break
	default:
		return "", false
	}

	return "", false
}

func compilerToExe(lines []string) {
	now := time.Now()
	unixNano := now.UnixNano()
	fmt.Println("Unix Timestamp (nanosaniye):", unixNano)

	outputfile := outputFileAddress + "/" + "c_" + strconv.FormatInt(unixNano, 10)
	fmt.Println(outputfile)
	f, _ := os.Create(outputfile + ".c")
	fmt.Println(f)
	kod := `#include <stdio.h>
int main() {
    printf("Hello, World!\n");
    return 0;
}
`
	_, _ = f.WriteString(kod)
	cmd := exec.Command("gcc", "./"+outputfile+".c", "-static", "-o", "./"+outputfile+".exe")
	fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Derleme hatası:", err)
		fmt.Println("gcc çıktısı:\n", string(out))
		return
	}

	fmt.Println("Derleme başarılı. Çıktı dosyası:", outputfile+".exe")
	fmt.Println("gcc çıktısı:\n", string(out))

}
