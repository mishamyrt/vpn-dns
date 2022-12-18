package network

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type Interface struct {
	Name string
}

// SetDNS sets interface domain name servers.
func (n *Interface) SetDNS(servers []string) (err error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("networksetup", "-setdnsservers", n.Name, strings.Join(servers, " "))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	if stderr.Len() > 0 {
		err = errors.New(stderr.String())
	}
	if stdout.Len() > 0 {
		err = errors.New(stdout.String())
	}
	return err
}
