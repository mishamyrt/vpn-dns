// Package vpn provides tools for working with the macOS system VPN.
package vpn

import (
	"net"
	"time"
	"vpn-dns/pkg/exec"
)

// ConnectionCheckInterval means the time that passes between connection checks.
var ConnectionCheckInterval = 700 * time.Millisecond

// Watcher is an entity that monitors changes in enabled VPNs and returns updates.
type Watcher struct {
	Names                   []string
	Updates                 chan *[]string
	Errors                  chan error
	CloseCheckInterval      time.Duration
	ConnectionCheckInterval time.Duration
	Execute                 exec.CommandRunner
	cmd                     exec.OSCommand
	closing                 bool
	closed                  bool
	inited                  bool
	statuses                []bool
	active                  []string
	IFacesCount             int
}

// Close watcher channels.
func (w *Watcher) Close() {
	w.closing = true
	for {
		if w.closed {
			break
		}
		time.Sleep(w.CloseCheckInterval)
	}
	close(w.Updates)
	close(w.Errors)
}

// Run watcher in goroutine.
func (w *Watcher) Run() {
	go w.start()
}

func (w *Watcher) start() {
	for {
		if w.closing {
			w.closed = true
			return
		}
		if !w.hasInterfaceChanges() {
			time.Sleep(w.ConnectionCheckInterval)
			continue
		}
		changed := w.collectState()
		if changed || !w.inited {
			w.inited = true
			w.Updates <- &w.active
		}
	}
}

func (w *Watcher) hasInterfaceChanges() bool {
	addrs, _ := net.InterfaceAddrs()
	ifacesCount := len(addrs)
	if ifacesCount != w.IFacesCount {
		w.IFacesCount = ifacesCount
		return true
	}
	return false
}

func (w *Watcher) collectState() bool {
	changed := false
	for i := range w.Names { //nolint:varnamelen
		status, err := IsConnected(w.Names[i], w.Execute)
		if err != nil {
			w.Errors <- err
			break
		}
		if status != w.statuses[i] {
			if !changed {
				changed = true
				w.active = nil
			}
			w.statuses[i] = status
			if status {
				w.active = append(w.active, w.Names[i])
			}
		}
	}
	return changed
}

// NewWatcher creates new VPN watcher.
func NewWatcher(names []string) Watcher {
	cmd := exec.OSCommand{}
	watcher := Watcher{
		Names:                   names,
		Updates:                 make(chan *[]string),
		Errors:                  make(chan error),
		ConnectionCheckInterval: ConnectionCheckInterval,
		CloseCheckInterval:      ConnectionCheckInterval,
		inited:                  false,
		cmd:                     cmd,
		Execute:                 cmd.Run,
		statuses:                make([]bool, len(names)),
		active:                  make([]string, 0),
	}
	return watcher
}
