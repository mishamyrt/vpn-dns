package network_test

import (
	"errors"
	"strings"
	"testing"
	"vpn-dns/pkg/exec"
	"vpn-dns/pkg/network"
)

const ifaceName = "Mock"

var nameServers = []string{"55.55.55.55", "11.11.11.11"}

func TestInterface(t *testing.T) {
	t.Parallel()
	mock := exec.Mock{}
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
	mock.Stderr.WriteString("Some error")
	err = iface.SetDNS([]string{"55.55.55.55", "11.11.11.11"})
	if !errors.Is(err, network.ErrDNSSet) {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.Clear()
	mock.ShoudFail = true
	err = iface.SetDNS([]string{"55.55.55.55", "11.11.11.11"})
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}
