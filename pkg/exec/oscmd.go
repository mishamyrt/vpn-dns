// Package exec provides functions to run external commands.
package exec

import (
	"bytes"
	goexec "os/exec"
)

// OSCommand represents OS command executor.
type OSCommand struct {
	Out bytes.Buffer
}

// Run command with given argument and return output.
func (c *OSCommand) Run(name string, args ...string) (string, error) {
	out := &c.Out
	cmd := goexec.Command(name, args...)
	cmd.Stdout = out
	cmd.Stderr = out
	err := cmd.Run()
	cmd = nil
	if out.Len() > 0 {
		result := out.String()
		c.Out.Reset()
		return result, err
	}
	return "", err
}
