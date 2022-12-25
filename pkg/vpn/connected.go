// Package vpn provides tools for working with the macOS system VPN.
package vpn

import (
	"errors"
	"strings"
	"vpn-dns/pkg/exec"
)

// ErrCommandFailed indicates that the VPN status could not be obtained.
var ErrCommandFailed = errors.New("scutil has been failed")

// IsConnected checks if connection with given name is active.
func IsConnected(name string, run exec.CommandRunner) (bool, error) {
	out, err := run("scutil", "--nc", "status", name)
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(out, "Connected"), nil
}
