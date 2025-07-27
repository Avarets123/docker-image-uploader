package common

import (
	"os"
	"os/exec"
)

func RunCmd(cmdStr string, args ...string) error {
	cmd := exec.Command(cmdStr, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()

}
