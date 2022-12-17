package app

type VPNConfig struct {
	Name    string   `yaml:"name"`
	Servers []string `yaml:"servers"`
}

type Config struct {
	Interface       string      `yaml:"interface"`
	VPNs            []VPNConfig `yaml:"VPNs"`
	FallbackServers []string    `yaml:"fallback_servers"`
}

func (c *Config) GetVPNs() []string {
	vpns := make([]string, len(c.VPNs))
	for i := range c.VPNs {
		vpns[i] = c.VPNs[i].Name
	}
	return vpns
}

func (c *Config) GetServers(activeVPNs []string) []string {
	servers := make([]string, 0)
	for i := range activeVPNs {
		vpnServers := c.findServers(activeVPNs[i])
		servers = append(servers, vpnServers...)
	}
	return servers
}

func (c *Config) findServers(vpn string) []string {
	for i := range c.VPNs {
		if c.VPNs[i].Name == vpn {
			return c.VPNs[i].Servers
		}
	}
	return []string{}
}
