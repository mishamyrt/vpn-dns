package app

import "errors"

// ErrFileNotExists is returned when requested file doesn't exists.
var ErrFileNotExists = errors.New("file does not exist")

// ErrDaemonNotRunning is returned when an inactive process is tried to perform an action.
var ErrDaemonNotRunning = errors.New("daemon is not running")
