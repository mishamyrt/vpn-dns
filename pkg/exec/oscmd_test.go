package exec_test

import (
	"strings"
	"testing"
	"vpn-dns/pkg/exec"
)

const testMsg = "Test message"

func assertOut(t *testing.T, cmd exec.OSCommand) {
	t.Helper()
	out, err := cmd.Run("echo", testMsg)
	if err != nil {
		t.Errorf("Unexpected error value: %v", err)
	}
	if strings.Trim(out, " \n") != testMsg {
		t.Errorf("Unexpected stdout value: %v", out)
	}
}

func assertErr(t *testing.T, cmd exec.OSCommand) {
	t.Helper()
	_, err := cmd.Run("exit", "1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	cmd := exec.OSCommand{}
	assertOut(t, cmd)
	assertErr(t, cmd)
}
