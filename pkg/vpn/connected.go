package vpn

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// IsConnected checks if connection with given name is active
func IsConnected(name string) (bool, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("scutil", "--nc", "status", name)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return false, err
	}
	if stderr.Len() > 0 {
		return false, errors.New(stderr.String())
	}
	return strings.HasPrefix(stdout.String(), "Connected"), nil
}
