package vpn

import (
	"time"
	"vpn-dns/pkg/exec"
)

// ConnectionCheckInterval means the time that passes between connection checks.
var ConnectionCheckInterval = 500 * time.Millisecond

type Watcher struct {
	Names                   []string
	Updates                 chan []string
	Errors                  chan error
	CloseCheckInterval      time.Duration
	ConnectionCheckInterval time.Duration
	_execute                exec.CommandRunner
	_closing                bool
	_closed                 bool
	_inited                 bool
}

func (w *Watcher) Close() {
	w._closing = true
	for {
		if w._closed {
			break
		}
		time.Sleep(w.CloseCheckInterval)
	}
	close(w.Updates)
	close(w.Errors)
}

func (w *Watcher) Run() {
	go w.start()
}

func (w *Watcher) start() {
	statuses := make([]bool, len(w.Names))
	var hasChanges bool
	var active []string
	for {
		if w._closing {
			w._closed = true
			return
		}
		hasChanges = false
		active = make([]string, 0)
		for i := range w.Names { //nolint:varnamelen
			status, err := IsConnected(w.Names[i], w._execute)
			if err != nil {
				w.Errors <- err
				break
			}
			if status != statuses[i] {
				hasChanges = true
				statuses[i] = status
				if status {
					active = append(active, w.Names[i])
				}
			}
		}
		if hasChanges || !w._inited {
			w._inited = true
			w.Updates <- active
		}
		time.Sleep(w.ConnectionCheckInterval)
	}
}

func NewWatcher(names []string, execute exec.CommandRunner) Watcher {
	watcher := Watcher{
		Names:                   names,
		Updates:                 make(chan []string),
		Errors:                  make(chan error),
		ConnectionCheckInterval: ConnectionCheckInterval,
		CloseCheckInterval:      ConnectionCheckInterval / 5,
		_inited:                 false,
		_execute:                execute,
	}
	return watcher
}
