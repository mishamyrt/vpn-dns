// Package exec provides functions to run external commands
package exec

import (
	"bytes"
	"errors"
	"strings"
)

// ErrMockCommand indicates that ShoudFail is active.
var ErrMockCommand = errors.New("mock error")

// MockCommand represents fake command executor, which output can be set.
// Simplifies testing.
type MockCommand struct {
	Out         bytes.Buffer
	LastCommand string
	ShoudFail   bool
}

// Run fake command. Returns the current values of buffers Stdout and Stderr.
// If ShoudFail flag is set to true, it returns ErrMockCommand error.
func (m *MockCommand) Run(name string, args ...string) (string, error) {
	m.LastCommand = name
	if len(args) > 0 {
		m.LastCommand += " " + strings.Join(args, " ")
	}
	var err error
	if m.ShoudFail {
		err = ErrMockCommand
	}
	result := m.Out.String()
	return result, err
}

// Clear output buffers.
func (m *MockCommand) Clear() {
	m.Out.Reset()
}
