package app

import (
	"io/ioutil"
	"os"
	"strconv"
)

func readPid() int {
	_, err := os.Stat(PidPath)
	if err != nil {
		return 0
	}
	data, err := ioutil.ReadFile(PidPath)
	if err != nil {
		return 0
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	return pid
}
