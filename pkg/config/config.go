package config

// Config represents vpn-dns configuration.
type Config struct {
	Interface       string
	FallbackServers []string
	VPNs            VPNs
}

// GetServers returns servers by active VPN connection names.
func (c *Config) GetServers(activeVPNs []string) ([]string, error) {
	servers := make([]string, 0)
	for i := range activeVPNs {
		vpnServers, err := c.VPNs.GetServers(activeVPNs[i])
		if err != nil {
			return servers, err
		}
		servers = append(servers, vpnServers...)
	}
	return servers, nil
}
