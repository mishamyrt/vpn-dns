// Package exec provides functions to run external commands.
package exec

import (
	"bytes"
	goexec "os/exec"
)

// CommandRunner represents function that can run external command and return output.
type CommandRunner func(name string, args ...string) (string, string, error)

// Run command with given argument and return output.
func Run(name string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := goexec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
