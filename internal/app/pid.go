package app

import (
	"os"
	"strconv"
)

func readPid() int {
	_, err := os.Stat(PidPath)
	if err != nil {
		return 0
	}
	data, err := os.ReadFile(PidPath)
	if err != nil {
		return 0
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	return pid
}
