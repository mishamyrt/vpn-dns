package vpn_test

import (
	"testing"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/vpn"
)

const connectionName = "connection_name"

func assertConnected(t *testing.T, expected bool, output string) {
	t.Helper()
	m := exec.Mock{}
	m.Stdout.WriteString("Disconnected\n")
	result, err := vpn.IsConnected(connectionName, m.Run)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if m.LastCommand != "scutil --nc status "+connectionName {
		t.Errorf("Unexpected command: %v", m.LastCommand)
	}
	if result {
		t.Errorf("Unexpected result: %v, should be %v", result, expected)
	}
}

func assertError(t *testing.T) {
	t.Helper()
	m := exec.Mock{}
	m.Stderr.WriteString("Bla bla bla, some error")
	_, err := vpn.IsConnected(connectionName, m.Run)
	if err != vpn.ErrCommandFailed {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestIsConnected(t *testing.T) {
	t.Parallel()
	assertConnected(t, true, "Connected\n")
	assertConnected(t, false, "Disconnected\n")
	assertError(t)
}
