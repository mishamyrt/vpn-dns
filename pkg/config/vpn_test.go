package config_test

import (
	"errors"
	"testing"
	"vpn-dns/pkg/config"
)

var MockVPNs = config.VPNs{
	"first":  []string{"1.1.1.1"},
	"second": []string{"2.2.2.2"},
	"third":  []string{"3.3.3.3"},
}

func TestNames(t *testing.T) {
	t.Parallel()
	names := MockVPNs.GetNames()
	if len(names) != 3 {
		t.Errorf("Unexpected names count: %d", len(names))
	}
}

func TestServers(t *testing.T) {
	t.Parallel()
	_, err := MockVPNs.GetServers("not_existing")
	if err == nil {
		t.Errorf("Unexpected nil")
	} else if !errors.Is(err, config.ErrNameNotFound) {
		t.Errorf("Unexpected error: %s", err.Error())
	}

	list, err := MockVPNs.GetServers("second")
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if len(list) != 1 || list[0] != "2.2.2.2" {
		t.Errorf("Unexpected results: %d", len(list))
	}
}
