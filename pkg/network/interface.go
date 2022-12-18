package network

import (
	"bytes"
	"os/exec"
	"strings"
)

type Interface struct {
	Name string
}

// SetDNS sets interface domain name servers.
func (n *Interface) SetDNS(servers []string) error {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("networksetup", "-setdnsservers", n.Name, strings.Join(servers, " "))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	if stderr.Len() > 0 || stdout.Len() > 0 {
		return ErrDNSSet
	}
	return nil
}
