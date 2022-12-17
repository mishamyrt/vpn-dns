package app

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readConfig(path string) (Config, error) {
	var configFile Config
	if fileExists(path) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return configFile, err
		}
		err = yaml.Unmarshal(data, &configFile)
		if err != nil {
			return configFile, err
		}
		return configFile, nil
	}
	return configFile, errors.New("file doesn't exist")
}
