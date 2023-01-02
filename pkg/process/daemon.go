// Package process provides support tools for dealing with background processes.
package process

import (
	"os"
	"strconv"
	"syscall"

	"github.com/sevlyar/go-daemon"
)

// Daemon represents background process.
type Daemon struct {
	PidPath string
	Context *daemon.Context
}

// Pid is current process id.
// If process is not running, returns 0.
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

// Running tells if process is running.
func (d *Daemon) Running() bool {
	pid := d.Pid()
	if pid == 0 {
		return false
	}
	return d.running(pid)
}

// Stop background process.
// Returns os.ErrProcessDone if process is not running.
func (d *Daemon) Stop() error {
	pid := d.Pid()
	if pid == 0 || !d.running(pid) {
		return os.ErrProcessDone
	}
	err := syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		return err
	}
	return os.Remove(d.PidPath)
}

func (d *Daemon) running(pid int) bool {
	process, _ := os.FindProcess(pid)
	err := process.Signal(syscall.Signal(0))
	return err == nil
}

// NewDaemon creates daemon instance.
func NewDaemon(name string) Daemon {
	pidPath := "/tmp/" + name + ".pid"
	logPath := "/tmp/" + name + ".log"
	return Daemon{
		PidPath: pidPath,
		Context: &daemon.Context{
			PidFileName: pidPath,
			LogFileName: logPath,
			WorkDir:     "./",
		},
	}
}
