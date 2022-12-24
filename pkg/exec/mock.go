package exec

import (
	"bytes"
	"errors"
	"strings"
)

type Mock struct {
	Stdout      bytes.Buffer
	Stderr      bytes.Buffer
	LastCommand string
	ShoudFail   bool
}

var ErrMockCommand = errors.New("mock error")

func (m *Mock) Run(name string, args ...string) (string, string, error) {
	m.LastCommand = name
	if len(args) > 0 {
		m.LastCommand += " " + strings.Join(args, " ")
	}
	var err error
	if m.ShoudFail {
		err = ErrMockCommand
	}
	return m.Stdout.String(), m.Stderr.String(), err
}

func (m *Mock) Clear() {
	m.Stdout.Reset()
	m.Stderr.Reset()
}
