// Package vpndns provides tools to switch DNS depending on VPN.
package vpndns

import (
	"log"
	"vpn-dns/pkg/config"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/network"
	"vpn-dns/pkg/vpn"
)

// Changer is an entity that monitors VPN changes and changes DNS.
type Changer struct {
	iface   network.Interface
	config  config.Config
	execute exec.CommandRunner
}

// Run watcher.
func (c *Changer) Run() {
	watcher := vpn.NewWatcher(
		c.config.VPNs.GetNames(),
		exec.Run,
	)
	watcher.Run()
	defer watcher.Close()
	for {
		select {
		case update := <-watcher.Updates:
			c.handleUpdate(update)
		case err := <-watcher.Errors:
			c.handleError(err)
		}
	}
}

func (c *Changer) handleUpdate(vpns []string) {
	if len(vpns) == 0 {
		log.Println("VPNs not connected, setting fallback servers")
		err := c.iface.SetDNS(c.config.FallbackServers)
		if err != nil {
			log.Println("Error while setting fallback DNS")
		}
	} else {
		servers, err := c.config.GetServers(vpns)
		if err != nil {
			c.handleError(err)
		}
		log.Println("Setting custom servers:", servers)
		err = c.iface.SetDNS(servers)
		if err != nil {
			c.handleError(err)
		}
	}
}

func (c *Changer) handleError(err error) {
	log.Println("Error:", err.Error())
}

// NewChanger creates new VPN DNS Changer.
// Returns error if error is unreadable.
func NewChanger(configPath string, run exec.CommandRunner) (Changer, error) {
	var changer Changer
	cfg, err := config.Read(configPath)
	if err != nil {
		return changer, err
	}
	changer.iface = network.NewInterface(cfg.Interface, run)
	changer.config = cfg
	changer.execute = run
	return changer, nil
}
