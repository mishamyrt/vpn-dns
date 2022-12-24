package login_test

import (
	"errors"
	"os"
	"strings"
	"testing"
	"vpn-dns/pkg/login"
)

const packageName = "testfile"
const command = "ls -la"

func assertWrite(t *testing.T, item login.Item) {
	t.Helper()
	err := item.Write()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !item.IsSet() {
		t.Errorf("File is not set")
	}
	file, err := os.ReadFile(item.Path)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.Contains(string(file), packageName) {
		t.Errorf("Unexpected content: %v", file)
	}
}

func assertRemove(t *testing.T, item login.Item) {
	t.Helper()
	err := item.Remove()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	_, err = os.Stat(item.Path)
	if !os.IsNotExist(err) {
		t.Errorf("File not removed, unexpected error: %v", err)
	}
	if item.IsSet() {
		t.Errorf("File is set")
	}
}

func assertErrors(t *testing.T, item login.Item) {
	t.Helper()
	err := item.Remove()
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Unexpected error: %v", err)
	}
	item.Path = "/etc/not_existing"
	if item.Write() == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestItem(t *testing.T) {
	t.Parallel()
	item := login.NewItem(packageName, command, os.Getenv("PWD")+"/_loginitem.prop")
	assertWrite(t, item)
	assertRemove(t, item)
	assertErrors(t, item)
}

func TestNewItem(t *testing.T) {
	t.Parallel()
	item := login.NewItem(packageName, command, "/")
	if item.PackageName != packageName {
		t.Errorf("Unexpected packageName: %v", item.PackageName)
	}
	if len(item.Command) != 2 || item.Command[0] != "ls" {
		t.Errorf("Unexpected command: %v", item.Command)
	}
}
