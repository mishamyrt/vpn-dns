package login_test

import (
	"os"
	"strings"
	"testing"
	"vpn-dns/internal/login"
)

const packageName = "testfile"
const command = "ls -la"

func TestItem(t *testing.T) {
	t.Parallel()
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := cwd + "/_loginitem.prop"
	item := login.NewItem(packageName, command, filePath)
	err = item.Write()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !item.IsSet() {
		t.Errorf("File is not set")
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !strings.Contains(string(file), packageName) {
		t.Errorf("Unexpected content: %v", file)
	}
	err = item.Remove()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		t.Errorf("File not removed, unexpected error: %v", err)
	}
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
