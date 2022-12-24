package network

import (
	"strings"
	"vpn-dns/pkg/exec"
)

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
		return ErrDNSSet
	}
	return nil
}

func NewInterface(name string, run exec.CommandRunner) Interface {
	return Interface{
		Name: name,
		run:  run,
	}
}
