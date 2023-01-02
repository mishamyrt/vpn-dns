package login_test

import (
	"os"
	"strings"
	"testing"
	"vpn-dns/pkg/login"
)

type PathBuilder func(packageName string) (string, error)

func assertPath(t *testing.T, build PathBuilder) {
	t.Helper()
	path, err := build(packageName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.HasSuffix(path, packageName+".plist") {
		t.Errorf("Expected path to ends with .plist extension: %v", path)
	}
}

func assertBrokenHome(t *testing.T, build PathBuilder) {
	t.Helper()
	os.Unsetenv("HOME")
	_, err := build(packageName)
	if err.Error() != "$HOME is not defined" {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestLaunchAgentPath(t *testing.T) {
	t.Parallel()
	assertPath(t, login.LaunchAgentPath)
	assertBrokenHome(t, login.LaunchAgentPath)
}
