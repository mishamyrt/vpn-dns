package network_test

import (
	"strings"
	"testing"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/network"
)

const ifaceName = "Mock"

var nameServers = []string{"55.55.55.55", "11.11.11.11"}

func TestInterface(t *testing.T) {
	t.Parallel()
	mock := exec.MockCommand{}
	iface := network.NewInterface(ifaceName, mock.Run)
	if iface.Name != ifaceName {
		t.Errorf("Unexpected name: %v", ifaceName)
	}
	err := iface.SetDNS(nameServers)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedCmd := "networksetup -setdnsservers Mock " + strings.Join(nameServers, " ")
	if mock.LastCommand != expectedCmd {
		t.Errorf("Unexpected command: %v", mock.LastCommand)
	}
	mock.Out.WriteString("Some error")
	err = iface.SetDNS(nameServers)
	if !exec.IsCommandErr(err) {
		t.Errorf("Unexpected error: %v", err)
	}
	mock.Clear()

	err = iface.SetDNS([]string{})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.HasSuffix(mock.LastCommand, "Empty") {
		t.Errorf("Unexpected command: '%v'", mock.LastCommand)
	}

	mock.Clear()
	mock.ShoudFail = true
	err = iface.SetDNS([]string{"55.55.55.55", "11.11.11.11"})
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}
