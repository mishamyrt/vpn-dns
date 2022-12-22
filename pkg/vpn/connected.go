package vpn

import (
	"errors"
	"strings"
	"vpn-dns/pkg/exec"
)

var ErrCommandFailed = errors.New("scutil has been failed")

// IsConnected checks if connection with given name is active.
func IsConnected(name string, run exec.CommandRunner) (bool, error) {
	stdout, stderr, err := run("scutil", "--nc", "status", name)
	if err != nil {
		return false, err
	}
	if len(stderr) > 0 {
		return false, ErrCommandFailed
	}
	return strings.HasPrefix(stdout, "Connected"), nil
}
