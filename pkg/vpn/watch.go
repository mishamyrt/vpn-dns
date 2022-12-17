package vpn

import "time"

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
		time.Sleep(100 * time.Millisecond)
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
		for i := range names {
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
		time.Sleep(500 * time.Millisecond)
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

// func getStatuses(names []string) ([]bool, err) {
// 	statuses := make([]bool, len(names))
// for i := range names {
// 	status, err := IsConnected(names[i])
// 	if err != nil {
// 		return statuses, err
// 	}
// 	statuses[i] = status
// }
// 	return statuses, nil
// }

// func Watch(names []string) {
// 	for {
// 		err, statuses := getStatuses(names)
// 		if err != nil {
// 			fmt.Println("Can't fill initial statuses")
// 		}
// 	}

// 	// Fill initial statuses
// 	for i := range names {
// 		status, err := IsConnected(names[i])
// 		if err != nil {

// 		}
// 		statuses[i] =
// 	}
// }
