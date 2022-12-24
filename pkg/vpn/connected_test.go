package vpn_test

import (
	"errors"
	"testing"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/vpn"
)

const connectionName = "connection_name"
const outConnected = "Connected\n"
const outDisconnected = "Disconnected\n"

func assertConnected(t *testing.T, expected bool, output string) {
	t.Helper()
	mock := exec.Mock{}
	mock.Stdout.WriteString(output)
	result, err := vpn.IsConnected(connectionName, mock.Run)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if mock.LastCommand != "scutil --nc status "+connectionName {
		t.Errorf("Unexpected command: %v", mock.LastCommand)
	}
	if result != expected {
		t.Errorf("Unexpected result: %v, should be %v", result, expected)
	}
}

func assertError(t *testing.T) {
	t.Helper()
	mock := exec.Mock{}
	mock.Stderr.WriteString("Bla bla bla, some error")
	_, err := vpn.IsConnected(connectionName, mock.Run)
	if !errors.Is(err, vpn.ErrCommandFailed) {
		t.Errorf("Unexpected error: %v", err)
	}
	mock.Clear()
	mock.ShoudFail = true
	_, err = vpn.IsConnected(connectionName, mock.Run)
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}

func TestIsConnected(t *testing.T) {
	t.Parallel()
	assertConnected(t, true, outConnected)
	assertConnected(t, false, outDisconnected)
	assertError(t)
}
