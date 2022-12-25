package vpn_test

import (
	"errors"
	"testing"
	"time"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/vpn"
)

func TestWatch(t *testing.T) {
	t.Parallel()
	mock := exec.MockCommand{}
	mock.Out.WriteString(outDisconnected)
	watcher := vpn.NewWatcher([]string{"first", "second"})
	watcher.Execute = mock.Run
	watcher.ConnectionCheckInterval = 10 * time.Millisecond
	watcher.Run()
	updatesCount := 0
	defer watcher.Close()
	for {
		updatesCount++
		select {
		case active := <-watcher.Updates:
			switch len(*active) {
			case 0:
				if updatesCount != 1 {
					t.Errorf("Got empty, expected value")
				}
				// Reset count to cause update
				watcher.IFacesCount = 0
				mock.Clear()
				mock.Out.WriteString(outConnected)
			case 2:
				if updatesCount != 2 {
					t.Errorf("Got values, expected to be empty")
				}
				watcher.IFacesCount = 0
				mock.Clear()
				mock.ShoudFail = true
			default:
				t.Errorf("Got unexpected value %v", "")
			}
		case err := <-watcher.Errors:
			if !errors.Is(err, exec.ErrMockCommand) {
				t.Errorf("Unexpected error %v", err)
			}
			return
		}
	}
}
