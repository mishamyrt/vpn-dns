package login_test

import (
	"strings"
	"testing"
	"vpn-dns/internal/login"
)

func TestLaunchAgentPath(t *testing.T) {
	t.Parallel()
	path, err := login.LaunchAgentPath(packageName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.HasSuffix(path, packageName+".plist") {
		t.Errorf("Unexpected path: %v", path)
	}
}
