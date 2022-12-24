package login_test

import (
	"log"
	"os"
	"strings"
	"testing"
	"vpn-dns/pkg/login"
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

	err = os.Unsetenv("HOME")
	if err != nil {
		log.Fatal(err)
	}
	_, err = login.LaunchAgentPath(packageName)
	if err == nil {
		t.Errorf("Unexpected nil")
	}
}
