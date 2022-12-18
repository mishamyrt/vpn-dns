package vpn

import "time"

// CloseCheckInterval means the time that passes between activity checks.
var CloseCheckInterval = 100 * time.Millisecond

// ConnectionCheckInterval means the time that passes between connection checks.
var ConnectionCheckInterval = 500 * time.Millisecond

type Watcher struct {
	Names       []string
	Updates     chan []string
	Errors      chan error
	shouldClose bool
	closed      bool
	inited      bool
}

func (w *Watcher) Close() {
	w.shouldClose = true
	for {
		if w.closed {
			break
		}
		time.Sleep(CloseCheckInterval)
	}
	close(w.Updates)
	close(w.Errors)
	w.closed = true
}

func (w *Watcher) run(names []string) {
	statuses := make([]bool, len(names))
	var hasChanges bool
	var active []string
	for {
		hasChanges = false
		active = make([]string, 0)
		for i := range names { //nolint:varnamelen
			status, err := IsConnected(names[i])
			if err != nil {
				w.Errors <- err
				break
			}
			if status != statuses[i] {
				hasChanges = true
				statuses[i] = status
				if status {
					active = append(active, names[i])
				}
			}
		}
		if hasChanges || !w.inited {
			w.inited = true
			w.Updates <- active
		}
		time.Sleep(ConnectionCheckInterval)
	}
}

func NewWatcher(names []string) Watcher {
	watcher := Watcher{
		Updates: make(chan []string),
		Errors:  make(chan error),
		inited:  false,
	}
	go watcher.run(names)
	return watcher
}
