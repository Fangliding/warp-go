package utils

import (
	"net/netip"
	"os/exec"
	"strconv"
)

func SetInterface(name string, ips ...string) (er error) {
	for _, ip := range ips {
		err := exec.Command("/sbin/ip", "addr", "add", ip, "dev", name).Run()
		if err != nil {
			er = err
		}
	}

	err := exec.Command("/sbin/ip", "link", "set", name, "up").Run()
	if err != nil {
		er = err
	}

	return
}

func SetRouteTable(name, endpoint string, fwmark uint32, dests ...netip.Prefix) (er error) {
	for _, dest := range dests {
		err := exec.Command("/sbin/ip", "route", "add", dest.String(), "dev", name, "table", "50000").Run()
		if err != nil {
			er = err
		}
	}

	if fwmark != 0 {
		err := exec.Command("/sbin/ip", "-4", "rule", "add", "not", "fwmark", strconv.Itoa(int(fwmark)), "table", "50000").Run()
		if err != nil {
			er = err
		}

		err = exec.Command("/sbin/ip", "-6", "rule", "add", "not", "fwmark", strconv.Itoa(int(fwmark)), "table", "50000").Run()
		if err != nil {
			er = err
		}
	}

	err := exec.Command("/sbin/ip", "-4", "rule", "add", "table", "main", "suppress_prefixlength", "0").Run()
	if err != nil {
		er = err
	}

	err = exec.Command("/sbin/ip", "-6", "rule", "add", "table", "main", "suppress_prefixlength", "0").Run()
	if err != nil {
		er = err
	}

	return
}

func RemoveRoute(endpoint string, fwmark uint32) (er error) {
	err := exec.Command("/sbin/ip", "route", "flush", "table", "50000").Run()
	if err != nil {
		er = err
	}

	if fwmark != 0 {
		err = exec.Command("/sbin/ip", "-6", "rule", "delete", "not", "fwmark", strconv.Itoa(int(fwmark)), "table", "50000").Run()
		if err != nil {
			er = err
		}

		err := exec.Command("/sbin/ip", "-4", "rule", "delete", "not", "fwmark", strconv.Itoa(int(fwmark)), "table", "50000").Run()
		if err != nil {
			er = err
		}
	}

	err = exec.Command("/sbin/ip", "-6", "rule", "delete", "table", "main", "suppress_prefixlength", "0").Run()
	if err != nil {
		er = err
	}

	err = exec.Command("/sbin/ip", "-4", "rule", "delete", "table", "main", "suppress_prefixlength", "0").Run()
	if err != nil {
		er = err
	}
	return
}
