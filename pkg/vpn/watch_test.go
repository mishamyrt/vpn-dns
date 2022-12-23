package vpn_test

import (
	"testing"
	"time"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/vpn"
)

func TestWatch(t *testing.T) {
	t.Parallel()
	mock := exec.Mock{}
	mock.Stdout.WriteString(outDisconnected)
	watcher := vpn.NewWatcher([]string{"first", "second"}, mock.Run)
	watcher.ConnectionCheckInterval = 10 * time.Millisecond
	watcher.Run()
	updatesCount := 0
	for active := range watcher.Updates {
		updatesCount++
		switch len(active) {
		case 0:
			if updatesCount != 1 {
				t.Errorf("Got empty, expected value")
			}
			mock.Clear()
			mock.Stdout.WriteString(outConnected)
		case 2:
			if updatesCount != 2 {
				t.Errorf("Got values, expected to be empty")
			}
			watcher.Close()
		default:
			t.Errorf("Got unexpected value %v", "")
		}
	}
}
