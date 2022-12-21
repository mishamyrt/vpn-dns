package config_test

import (
	"errors"
	"strings"
	"testing"
	"vpn-dns/pkg/config"
)

const configPath = "../../testdata/basic-config.yaml"
const packagePath = "../../pkg"
const mainPath = "../../main.go"
const notExistingPath = "nanananananananana/BATMAN"

func assertFileNotExists(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected ErrFileNotExists, got nil")
	} else if !errors.Is(err, config.ErrFileNotExists) {
		t.Errorf("Unexpected error value: %v", err.Error())
	}
}

func TestRead(t *testing.T) {
	t.Parallel()
	// Check not existing path
	_, err := config.Read(notExistingPath)
	assertFileNotExists(t, err)
	// Check folder path
	_, err = config.Read(packagePath)
	assertFileNotExists(t, err)
	// Check not correct file
	_, err = config.Read(mainPath)
	errMessage := err.Error()
	if !strings.Contains(errMessage, "yaml:") {
		t.Errorf("Expected yaml error, got %v", errMessage)
	}
	// Check correct file
	cfg, err := config.Read(configPath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	if cfg.Interface != "Wi-Fi" {
		t.Errorf("Unexpected config interface value: %v", cfg.Interface)
	}
}
