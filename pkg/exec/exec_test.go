package exec_test

import (
	"strings"
	"testing"
	"vpn-dns/pkg/exec"
)

const testMsg = "Test message"

func assertStdout(t *testing.T, run exec.CommandRunner) {
	t.Helper()
	stdout, stderr, err := run("echo", testMsg)
	if err != nil {
		t.Errorf("Unexpected error value: %v", stderr)
	}
	if strings.Trim(stdout, " \n") != testMsg {
		t.Errorf("Unexpected stdout value: %v", stdout)
	}
	if len(stderr) > 0 {
		t.Errorf("Unexpected stderr value: %v", stderr)
	}
}

func assertErr(t *testing.T, run exec.CommandRunner) {
	t.Helper()
	_, _, err := run("exit", "1")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func assertStderr(t *testing.T, run exec.CommandRunner) {
	t.Helper()
	stdout, stderr, _ := run("ls", "not_existing")
	if len(stderr) == 0 {
		t.Errorf("Unexpected empty stderr")
	}
	if len(stdout) > 0 {
		t.Errorf("Unexpected stdout value: %v", stdout)
	}
}

func TestRun(t *testing.T) {
	t.Parallel()
	assertStdout(t, exec.Run)
	assertErr(t, exec.Run)
	assertStderr(t, exec.Run)
}
