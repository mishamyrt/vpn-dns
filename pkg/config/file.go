package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// VPNConfig represents YAML configuration VPN entry.
type VPNEntry struct {
	Name    string   `yaml:"name"`
	Servers []string `yaml:"servers"`
}

// ConfigFile represents YAML configuration.
type File struct {
	Interface       string     `yaml:"interface"`
	VPNs            []VPNEntry `yaml:"VPNs"`
	FallbackServers []string   `yaml:"fallback_servers"`
}

// Exists checks if file is exists.
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// ErrFileNotExists is returned when requested file doesn't exists.
var ErrFileNotExists = errors.New("configuration file doesn't exist")

// Read configuration file.
func Read(path string) (Config, error) {
	var configFile File
	var config Config
	if !Exists(path) {
		return config, ErrFileNotExists
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
