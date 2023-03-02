package utils

import (
	"errors"

	"golang.zx2c4.com/wireguard/conn"
)

type Bind struct {
	conn.Bind
	reseved [3]byte
}

func NewResevedBind(bind conn.Bind, reserved [3]byte) *Bind {
	return &Bind{
		Bind:    bind,
		reseved: reserved,
	}
}

func (b *Bind) SetReseved(reserved [3]byte) {
	b.reseved = reserved
}

func (b *Bind) Send(buf []byte, ep conn.Endpoint) error {
	if len(buf) > 3 {
		buf[1] = b.reseved[0]
		buf[2] = b.reseved[1]
		buf[3] = b.reseved[2]
	}

	return b.Bind.Send(buf, ep)
}

func (b *Bind) Open(port uint16) (fns []conn.ReceiveFunc, actualPort uint16, err error) {
	fns, actualPort, err = b.Bind.Open(port)
	if err != nil {
		return
	}

	var tempFns []conn.ReceiveFunc
	for _, fn := range fns {
		tempFns = append(tempFns,b.NewReceiveFunc(fn))
	}

	fns = tempFns
	return
}

func (b *Bind) NewReceiveFunc(fn conn.ReceiveFunc) conn.ReceiveFunc {
	return func(buf []byte) (n int, ep conn.Endpoint, err error) {
		n, ep, err = fn(buf)
		if err != nil || n < 4 {
			return
		}

		if buf[1] != b.reseved[0] || buf[2] != b.reseved[1] || buf[3] != b.reseved[2] {
			err = errors.New("bad reseved")
			return
		}

		buf[1] = 0
		buf[2] = 0
		buf[3] = 0
		return
	}
}
