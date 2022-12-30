package exec_test

import (
	"errors"
	"strings"
	"testing"
	"vpn-dns/pkg/exec"
)

const errMessage = "command failed"

func TestError(t *testing.T) {
	t.Parallel()
	cmdErr := exec.NewCommandErr(errMessage)
	if !strings.HasSuffix(cmdErr.Error(), errMessage) {
		t.Errorf("Unexpected error: %v", cmdErr)
	}

	if !exec.IsCommandErr(cmdErr) {
		t.Errorf("Expected to be CommandErr")
	}

	notCmdErr := errors.New("some error") //nolint: goerr113
	if exec.IsCommandErr(notCmdErr) {
		t.Errorf("Expected not to be CommandErr")
	}
}
