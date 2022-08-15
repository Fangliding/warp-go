package utils

import (
	"errors"
	"net"
	"net/netip"
	"os/exec"
)

func SetInterface(name string, ips ...string) (er error) {
	err := exec.Command("/sbin/ifconfig", name, "up").Run()
	if err != nil {
		er = err
	}

	for _, ip := range ips {
		if net.ParseIP(ip).To16() == nil {
			err = exec.Command("/sbin/ifconfig", name, "inet", "add", ip+"/32").Run()
		} else {
			err = exec.Command("/sbin/ifconfig", name, "inet6", "add", ip+"/128").Run()
		}
		if err != nil {
			er = err
		}
	}

	return
}

func SetRouteTable(name, endpoint string, fwmark uint32, dests ...netip.Prefix) (er error) {
	gw, err := exec.Command("/usr/bin/csh", "-c", "netstat -r | grep default | head -n 2 | awk '{print $2}'").Output()
	if err != nil {
		er = err
		return
	}

	if string(gw) == "" {
		er = errors.New("failed to get gateway")
		return
	}

	for _, dest := range dests {
		err := exec.Command("/sbin/route", "-q", "-n", "add", dest.String(), "-interface", name).Run()
		if err != nil {
			er = err
		}
	}

	err = exec.Command("/sbin/route", "-q", "-n", "delete", endpoint).Run()
	if err != nil {
		er = err
	}

	err = exec.Command("/sbin/route", "-q", "-n", "add", endpoint, "-gateway", string(gw)).Run()
	if err != nil {
		er = err
	}

	return
}

func RemoveRoute(endpoint string, fwmark uint32) (er error) {
	err := exec.Command("/sbin/route", "-q", "-n", "delete", endpoint).Run()
	if err != nil {
		er = err
	}

	return
}
