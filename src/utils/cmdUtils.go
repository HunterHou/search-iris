package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

func ExecCmdStart(path string) {
	ExecCmd(path,"start")
}

func ExecCmdExplorer(path string) {
	ExecCmd(path,"explorer")
}

func ExecCmd(path string,cmdType string) {
	cmd := exec.Command("cmd", "/C",cmdType, "", path)
	if cmd != nil {
		fmt.Println(cmd.String())
		if runtime.GOOS == "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		}
		cmdErr := cmd.Start()
		if cmdErr != nil {
			fmt.Println(cmdErr)
		}
	}
}
