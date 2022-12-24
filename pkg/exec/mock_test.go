package exec_test

import (
	"testing"
	"vpn-dns/pkg/exec"
)

const command = "command"
const stdoutMsg = "out-message"
const stderrMsg = "err-message"

func TestMock(t *testing.T) {
	t.Parallel()
	mock := exec.Mock{}
	mock.Stdout.WriteString(stdoutMsg)
	mock.Stderr.WriteString(stderrMsg)

	stdout, stderr, _ := mock.Run(command, "1", "2", "3")
	if stdout != stdoutMsg {
		t.Errorf("Unexpected stdout value: %v", stdout)
	}
	if stderr != stderrMsg {
		t.Errorf("Unexpected stderr value: %v", stderr)
	}
	if mock.LastCommand != command+" 1 2 3" {
		t.Errorf("Unexpected command: %v", mock.LastCommand)
	}

	mock.Clear()
	stdout, stderr, _ = mock.Run(command)
	if len(stdout) > 0 || len(stderr) > 0 {
		t.Errorf("Unexpected values: out - %v; err - %v", stdout, stderr)
	}

	mock.ShoudFail = true
	_, _, err := mock.Run(command)
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}
