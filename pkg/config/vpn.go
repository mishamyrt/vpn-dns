package config

import (
	"os"
)

// VPNs represents names to servers mapping.
type VPNs map[string][]string

// GetServers returns servers by name.
func (v *VPNs) GetServers(name string) ([]string, error) {
	if servers, found := map[string][]string(*v)[name]; found {
		return servers, nil
	}
	return []string{}, os.ErrNotExist
}

// GetNames returns all VPN names in configuration.
func (v *VPNs) GetNames() []string {
	names := make([]string, len(*v))
	i := 0
	for key := range map[string][]string(*v) {
		names[i] = key
		i++
	}
	return names
}
