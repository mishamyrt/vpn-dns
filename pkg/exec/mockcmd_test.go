package exec_test

import (
	"testing"
	"vpn-dns/pkg/exec"
)

const command = "command"
const outMsg = "out-message"

func TestMock(t *testing.T) {
	t.Parallel()
	mock := exec.MockCommand{}
	mock.Out.WriteString(outMsg)

	out, _ := mock.Run(command, "1", "2", "3")
	if out != outMsg {
		t.Errorf("Unexpected stdout value: %v", out)
	}
	if mock.LastCommand != command+" 1 2 3" {
		t.Errorf("Unexpected command: %v", mock.LastCommand)
	}

	mock.Clear()
	out, _ = mock.Run(command)
	if len(out) > 0 {
		t.Errorf("Unexpected value: %v", out)
	}

	mock.ShoudFail = true
	_, err := mock.Run(command)
	if err == nil {
		t.Errorf("Unexpected nil error")
	}
}
