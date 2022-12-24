package process

import (
	"os"
	"strconv"
	"syscall"

	"github.com/sevlyar/go-daemon"
)

type Daemon struct {
	PidPath string
	LogPath string
}

func (d *Daemon) Pid() int {
	_, err := os.Stat(d.PidPath)
	if err != nil {
		return 0
	}
	pidData, err := os.ReadFile(d.PidPath)
	if err != nil {
		return 0
	}
	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		return 0
	}
	return pid
}

func (d *Daemon) Start() error {
	if d.Running() {
		return os.ErrExist
	}
	cntxt := &daemon.Context{
		PidFileName: d.PidPath,
		LogFileName: d.LogPath,
		WorkDir:     "./",
	}

	child, err := cntxt.Reborn()
	if err != nil {
		return err
	}
	if child != nil {
		return ErrChildNotDone
	}
	defer cntxt.Release()

	return nil
}

func (d *Daemon) Running() bool {
	pid := d.Pid()
	if pid == 0 {
		return false
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (d *Daemon) Stop() error {
	pid := d.Pid()
	if pid == 0 {
		return os.ErrProcessDone
	}
	return syscall.Kill(pid, syscall.SIGINT)
}

func NewDaemon(name string) Daemon {
	return Daemon{
		PidPath: "/tmp/" + name + ".pid",
		LogPath: "/tmp/" + name + ".log",
	}
}
