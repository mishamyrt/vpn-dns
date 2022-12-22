package login

import (
	"os"
	"strings"
)

type Item struct {
	PackageName string
	Command     []string
	path        string
}

func (it *Item) IsSet() bool {
	info, err := os.Stat(it.path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (it *Item) Remove() error {
	if it.IsSet() {
		return os.Remove(it.path)
	}
	return nil
}

func (it *Item) Write() error {
	content := it.Render()
	err := os.WriteFile(it.path, []byte(content), 0644) //nolint:gomnd
	if err != nil {
		return err
	}
	return nil
}

func (it *Item) Render() string {
	list := NewPropList()
	list.Bool("KeepAlive", false)
	list.String("Label", it.PackageName)
	list.StringArray("ProgramArguments", it.Command)
	list.Bool("RunAtLoad", true)
	list.String("StandardErrorPath", "/dev/null")
	list.String("StandardOutPath", "/dev/null")
	return list.Join()
}

func NewItem(name string, command string, path string) Item {
	item := Item{
		PackageName: name,
		Command:     strings.Split(command, " "),
		path:        path,
	}
	return item
}

func LaunchAgentPath(name string) (string, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return userDir + "/Library/LaunchAgents/" + name + ".plist", nil
}
