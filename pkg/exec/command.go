// Package exec provides functions to run external commands.
package exec

// CommandRunner represents function that can run external command and return output.
type CommandRunner func(name string, args ...string) (string, error)
