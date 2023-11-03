package gou

import (
	"os/exec"
	"strings"
)

func RunCmd(cmdStr string) (error, string) {
	out := new(strings.Builder)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = out
	return cmd.Run(), out.String()
}

func CmdExists(cmdStr string) bool {
	_, err := exec.LookPath(cmdStr)
	if err != nil {
		return false
	}
	return true
}
