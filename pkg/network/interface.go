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
	var dnsConfig string
	if len(servers) > 0 {
		dnsConfig = strings.Join(servers, " ")
	} else {
		dnsConfig = "Empty"
	}

	out, err := n.run("networksetup", "-setdnsservers", n.Name, dnsConfig)
	if err != nil || len(out) > 0 {
		return exec.NewCommandErr(out)
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
