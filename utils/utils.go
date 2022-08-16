package utils

import (
	"fmt"
	"os/exec"
)

func Exec(cmds ...string) (e error) {
	for _, cmd := range cmds {
		err := exec.Command(SHELL, "-c", cmd).Run()
		if err != nil {
			e = err
		}
	}

	return
}
