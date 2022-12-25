// Package exec provides functions to run external commands
package exec

import (
	"bytes"
	"errors"
	"strings"
)

// ErrMockCommand indicates that ShoudFail is active.
var ErrMockCommand = errors.New("mock error")

// Mock represents fake executor, which output can be set.
// Simplifies testing.
type Mock struct {
	Stdout      bytes.Buffer
	Stderr      bytes.Buffer
	LastCommand string
	ShoudFail   bool
}

// Run fake command. Returns the current values of buffers Stdout and Stderr.
// If ShoudFail flag is set to true, it returns ErrMockCommand error.
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

// Clear output buffers.
func (m *Mock) Clear() {
	m.Stdout.Reset()
	m.Stderr.Reset()
}
