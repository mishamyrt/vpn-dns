package login

import (
	"os"
	"strings"
)

type Item struct {
	PackageName string
	Command     []string
	Path        string
}

func (it *Item) IsSet() bool {
	info, err := os.Stat(it.Path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (it *Item) Remove() error {
	if it.IsSet() {
		return os.Remove(it.Path)
	}
	return os.ErrNotExist
}

func (it *Item) Write() error {
	content := it.render()
	err := os.WriteFile(it.Path, []byte(content), 0644) //nolint:gomnd
	if err != nil {
		return err
	}
	return nil
}

func (it *Item) render() string {
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
		Path:        path,
	}
	return item
}
