// Package config provides tools for working with the VPN DNS configuration file.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// VPNEntry represents YAML configuration VPN entry.
type VPNEntry struct {
	Name    string   `yaml:"name"`
	Servers []string `yaml:"servers"`
}

// File represents YAML configuration.
type File struct {
	Interface       string     `yaml:"interface"`
	VPNs            []VPNEntry `yaml:"VPNs"`
	FallbackServers []string   `yaml:"fallback_servers,omitempty"`
}

// Read configuration file.
// Returns error if file not exist or unreadable.
func Read(path string) (Config, error) {
	var configFile File
	var config Config
	if !fileExists(path) {
		return config, os.ErrNotExist
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &configFile)
	if err != nil {
		return config, err
	}
	config.Interface = configFile.Interface
	config.FallbackServers = configFile.FallbackServers
	config.VPNs = make(VPNs)
	for _, entry := range configFile.VPNs {
		config.VPNs[entry.Name] = entry.Servers
	}
	return config, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
