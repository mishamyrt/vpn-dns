package config_test

import (
	"errors"
	"testing"
	"vpn-dns/pkg/config"
)

func TestGetServers(t *testing.T) {
	t.Parallel()
	cfg := config.Config{
		VPNs: MockVPNs,
	}
	single, err := cfg.GetServers([]string{"first"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if single[0] != "1.1.1.1" {
		t.Errorf("Unexpected output: %v", single[0])
	}
	multiple, err := cfg.GetServers([]string{"first", "third"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if multiple[0] != "1.1.1.1" || multiple[1] != "3.3.3.3" {
		t.Errorf("Unexpected output: 1. %v;  2. %v", multiple[0], multiple[1])
	}
	_, err = cfg.GetServers([]string{"last"})
	if !errors.Is(err, config.ErrNotExist) {
		t.Errorf("Unexpected error: %v", err)
	}
}
