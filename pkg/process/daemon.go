package process

import (
	"os"
	"strconv"
	"syscall"

	"github.com/sevlyar/go-daemon"
)

type Daemon struct {
	PidPath string
	Context *daemon.Context
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

func (c *Daemon) Running() bool {
	pid := c.Pid()
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
	pidPath := "/tmp/" + name + ".pid"
	logPath := "/tmp/" + name + ".log"
	return Daemon{
		PidPath: pidPath,
		Context: &daemon.Context{
			PidFileName: pidPath,
			LogFileName: logPath,
			PidFilePerm: 0644, //nolint:gomnd
			WorkDir:     "./",
		},
	}
}
