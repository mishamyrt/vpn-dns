// Package exec provides functions to run external commands.
package exec

import (
	"errors"
	"strings"
)

const commandErrPrefix = "command error: "

// NewCommandErr creates new command error.
func NewCommandErr(output string) error {
	return errors.New(commandErrPrefix + output) //nolint: goerr113
}

// IsCommandErr checks if error is caused by command fail.
func IsCommandErr(err error) bool {
	return strings.HasPrefix(err.Error(), commandErrPrefix)
}
