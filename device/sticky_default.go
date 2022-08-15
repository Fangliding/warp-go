//go:build !linux

package device

import (
	"warp-go/conn"
	"warp-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
