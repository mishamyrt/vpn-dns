package process_test

import (
	"errors"
	"os"
	"testing"
	"time"
	"vpn-dns/pkg/process"
)

const appName = "test-app"

func assertParent(t *testing.T, daemon process.Daemon) {
	t.Helper()
	if !daemon.Running() {
		t.Errorf("Expected daemon to be running")
	}
	err := daemon.Stop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if daemon.Running() {
		t.Errorf("Expected daemon to be stopped")
	}
	pid := daemon.Pid()
	if pid != 0 {
		t.Errorf("Unexpected process id value: %v", pid)
	}
	err = daemon.Stop()
	if !errors.Is(err, os.ErrProcessDone) {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestDaemon(t *testing.T) {
	t.Parallel()
	daemon := process.NewDaemon(appName)
	child, err := daemon.Context.Reborn()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if child != nil {
		time.Sleep(50 * time.Millisecond)
		assertParent(t, daemon)
		return
	}
	defer daemon.Context.Release() //nolint: errcheck
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
