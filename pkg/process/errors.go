package process

import "errors"

var ErrChildNotDone = errors.New("child process isn't finished")
