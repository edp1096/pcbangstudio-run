package main // import "run"

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// cmd := exec.Command(dir + "/tools/vscode/run_main.cmd")
	cmd := exec.Command(dir+"/tools/vscode/RunHiddenConsole.exe", "run_vscode.cmd")
	cmd.Dir = dir + "/tools/vscode"

	stdout, err := cmd.StdoutPipe()
	err = cmd.Start()
	if err != nil {
		log.Printf("Error occured: %v", err)
	}

	bResult, _ := ioutil.ReadAll(stdout)
	fmt.Println(string(bResult))
}
