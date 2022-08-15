package utils

import "os"

var SHELL = "/usr/bin/bash"

func init() {
	if _, err := os.Stat(SHELL); err != nil {
		SHELL = "/usr/bin/zsh"
	}
}
