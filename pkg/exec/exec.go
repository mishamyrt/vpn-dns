package exec

import (
	"bytes"
	goexec "os/exec"
	"strings"
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

type Mock struct {
	Stdout      bytes.Buffer
	Stderr      bytes.Buffer
	LastCommand string
}

func (m *Mock) Run(name string, args ...string) (string, string, error) {
	m.LastCommand = name + " " + strings.Join(args, " ")
	return m.Stdout.String(), m.Stderr.String(), nil
}

func (m *Mock) Clear(name string, args ...string) (string, string, error) {
	m.Stdout.Reset()
	m.Stderr.Reset()
	return m.Stdout.String(), m.Stderr.String(), nil
}
