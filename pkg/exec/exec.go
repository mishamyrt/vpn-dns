package exec

import (
	"bytes"
	goexec "os/exec"
)

type CommandRunner func(name string, args ...string) (string, string, error)

func Run(name string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := goexec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
