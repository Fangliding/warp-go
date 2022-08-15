package utils

import (
	"net"
	"os/exec"
)

func SetInterface(name string, ips ...string) (er error) {
	for _, ip := range ips {
		var err error
		if net.ParseIP(ip).To16() == nil {
			err = exec.Command("netsh.exe", "interface", "ipv4", "add", "address", name, ip).Run()
		} else {
			err = exec.Command("netsh.exe", "interface", "ipv6", "add", "address", name, ip).Run()
		}

		if err != nil {
			er = err
		}
	}

	return
}
