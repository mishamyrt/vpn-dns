// Package vpndns provides tools to switch DNS depending on VPN.
package vpndns

import (
	"log"
	"os"
	"strings"
	"vpn-dns/pkg/config"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/network"
	"vpn-dns/pkg/vpn"
)

// Changer is an entity that monitors VPN changes and changes DNS.
type Changer struct {
	iface   network.Interface
	config  config.Config
	cmd     exec.OSCommand
	watcher vpn.Watcher
}

// Run watcher.
func (c *Changer) Run() {
	c.watcher = vpn.NewWatcher(c.config.VPNs.GetNames())
	c.watcher.Run()
	defer c.watcher.Close()
	for {
		select {
		case update := <-c.watcher.Updates:
			c.handleUpdate(update)
		case err := <-c.watcher.Errors:
			c.handleError(err)
		}
	}
}

func (c *Changer) handleUpdate(vpns *[]string) {
	if len(*vpns) == 0 {
		log.Println("VPNs not connected, setting fallback servers")
		err := c.iface.SetDNS(c.config.FallbackServers)
		if err != nil {
			c.handleError(err)
		}
	} else {
		servers, err := c.config.GetServers(*vpns)
		if err != nil {
			c.handleError(err)
		}
		log.Printf("Setting custom servers: %v", servers)
		err = c.iface.SetDNS(servers)
		if err != nil {
			c.handleError(err)
		}
	}
}

func (c *Changer) handleError(err error) {
	if strings.Contains(err.Error(), "signal: killed") {
		log.Println("Got sigkill, exiting.")
		c.watcher.Close()
		os.Exit(0)
	}
	log.Printf("Error: %v", err)
}

// NewChanger creates new VPN DNS Changer.
// Returns error if error is unreadable.
func NewChanger(configPath string) (Changer, error) {
	var changer Changer
	cmd := exec.OSCommand{}
	cfg, err := config.Read(configPath)
	if err != nil {
		return changer, err
	}
	changer.iface = network.NewInterface(cfg.Interface, cmd.Run)
	changer.config = cfg
	changer.cmd = cmd
	return changer, nil
}
