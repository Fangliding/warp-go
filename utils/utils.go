package utils

import (
	"os/exec"
)

func Exec(cmds ...string) {
	for _, cmd := range cmds {
		exec.Command(SHELL, "-c", cmd).Run()
	}
}
