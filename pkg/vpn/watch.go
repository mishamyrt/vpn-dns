// Package vpn provides tools for working with the macOS system VPN.
package vpn

import (
	"time"
	"vpn-dns/pkg/exec"
)

// ConnectionCheckInterval means the time that passes between connection checks.
var ConnectionCheckInterval = 500 * time.Millisecond

// Watcher is an entity that monitors changes in enabled VPNs and returns updates.
type Watcher struct {
	Names                   []string
	Updates                 chan []string
	Errors                  chan error
	CloseCheckInterval      time.Duration
	ConnectionCheckInterval time.Duration
	execute                 exec.CommandRunner
	closing                 bool
	closed                  bool
	inited                  bool
	statuses                []bool
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
		connected, changed := w.collectState()
		if changed || !w.inited {
			w.inited = true
			w.Updates <- connected
		}
		time.Sleep(w.ConnectionCheckInterval)
	}
}

func (w *Watcher) collectState() ([]string, bool) {
	connected := make([]string, 0)
	changed := false
	for i := range w.Names { //nolint:varnamelen
		status, err := IsConnected(w.Names[i], w.execute)
		if err != nil {
			w.Errors <- err
			break
		}
		if status != w.statuses[i] {
			changed = true
			w.statuses[i] = status
			if status {
				connected = append(connected, w.Names[i])
			}
		}
	}
	return connected, changed
}

// NewWatcher creates new VPN watcher.
func NewWatcher(names []string, execute exec.CommandRunner) Watcher {
	watcher := Watcher{
		Names:                   names,
		Updates:                 make(chan []string),
		Errors:                  make(chan error),
		ConnectionCheckInterval: ConnectionCheckInterval,
		CloseCheckInterval:      ConnectionCheckInterval,
		inited:                  false,
		execute:                 execute,
		statuses:                make([]bool, len(names)),
	}
	return watcher
}
