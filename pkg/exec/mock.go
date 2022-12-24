package exec

import (
	"bytes"
	"strings"
)

type Mock struct {
	Stdout      bytes.Buffer
	Stderr      bytes.Buffer
	LastCommand string
}

func (m *Mock) Run(name string, args ...string) (string, string, error) {
	m.LastCommand = name + " " + strings.Join(args, " ")
	return m.Stdout.String(), m.Stderr.String(), nil
}

func (m *Mock) Clear() {
	m.Stdout.Reset()
	m.Stderr.Reset()
}
