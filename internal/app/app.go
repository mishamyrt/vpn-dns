package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"

	"vpn-dns/pkg/network"
	"vpn-dns/pkg/vpn"
)

const (
	PidPath = "/tmp/vpn-dns.pid"
	LogPath = "/tmp/vpn-dns.log"
)

type App struct {
	pid    int
	config Config
}

func (a *App) Run() {
	iface := network.Interface{
		Name: a.config.Interface,
	}
	vpns := a.config.GetVPNs()
	watcher := vpn.NewWatcher(vpns)
	defer watcher.Close()
	for {
		select {
		case active, ok := <-watcher.Updates:
			if !ok {
				return
			}
			if len(active) == 0 {
				log.Println("VPNs not connected, setting fallback servers")
				iface.SetDNS(a.config.FallbackServers)
			} else {
				servers := a.config.GetServers(active)
				log.Println("Setting custom servers:", servers)
				iface.SetDNS(servers)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}

func (a *App) Running() bool {
	if a.pid == 0 {
		return false
	}
	process, err := os.FindProcess(a.pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (a *App) Kill() error {
	if a.pid == 0 {
		return errors.New("Daemon is not running")
	}
	return syscall.Kill(a.pid, syscall.SIGINT)
}

func Create(configPath string) App {
	config, err := readConfig(configPath)
	if err != nil {
		fmt.Println("Error while reading config:", err.Error())
	}
	return App{
		config: config,
		pid:    readPid(),
	}
}
