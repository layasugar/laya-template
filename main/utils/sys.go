package utils

import (
	"os"
	"os/exec"
)

// 守护执行
func Daemon() {
	args := os.Args
	var narcs []string
	for _, arg := range args {
		d := arg == "-d"
		daemon := arg == "-daemon"
		if d && daemon {
			narcs = append(narcs, arg)
		}
	}
	cmd := exec.Command(args[0], narcs...)
	_ = cmd.Start()
	os.Exit(0)
}
