package gcnf

import (
	"os"
	"strconv"
)

var (
	pid       int
	pidString string
)

// PID 得到 PID
func PID() int {
	if pid != 0 {
		return pid
	}

	pid = os.Getpid()
	pidString = strconv.Itoa(pid)
	return pid
}

// PIDString 得到PID 字符串形式
func PIDString() string {
	if pidString == "" {
		PID()
	}

	return pidString
}
