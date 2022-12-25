// Package network provides support tools for working with network devices.
package network

import (
	"strings"
	"vpn-dns/pkg/exec"
)

// Interface represents network interface controller.
type Interface struct {
	Name string
	run  exec.CommandRunner
}

// SetDNS sets interface domain name servers.
func (n *Interface) SetDNS(servers []string) error {
	dnsConfig := strings.Join(servers, " ")
	stdout, stderr, err := n.run("networksetup", "-setdnsservers", n.Name, dnsConfig)
	if err != nil {
		return err
	}
	if len(stderr) > 0 || len(stdout) > 0 {
		return ErrNotSet
	}
	return nil
}

// NewInterface creates new network interface controller entity.
func NewInterface(name string, run exec.CommandRunner) Interface {
	return Interface{
		Name: name,
		run:  run,
	}
}
