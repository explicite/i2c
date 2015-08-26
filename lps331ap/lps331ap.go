package lps331ap

import (
	"errors"

	"github.com/explicite/i2c"
)

type LPS331AP struct {
	bus    *i2c.Bus
	addr   byte
	active bool
}

func (l *LPS331AP) Init(addr byte, bus byte) error {
	var err error
	l.addr = addr
	l.bus, err = i2c.NewBus(bus)
	return err
}

func (l *LPS331AP) Read(reg byte) (byte, error) {
	buf, err := l.bus.Read(l.addr, reg, 1)
	if err != nil {
		return 0, err
	}

	return buf[0], nil
}

func (l *LPS331AP) Pressure() (float32, error) {
	buf := make([]byte, 3)

	for idx := 0x28; idx <= 0x2a; idx++ {
		var err error
		buf[idx-0x28], err = l.Read(byte(idx))
		if err != nil {
			return 0, err
		}
	}

	return float32(int(buf[2])<<16|int(buf[1])<<8|int(buf[0])) / 4096.0, nil
}

func (l *LPS331AP) Temperature() (float32, error) {
	buf := make([]byte, 2)

	for idx := 0x2b; idx <= 0x2c; idx++ {
		var err error
		buf[idx-0x2b], err = l.Read(byte(idx))
		if err != nil {
			return 0, err
		}
	}

	return 42.5 + float32(^(int16(buf[1])<<8|int16(buf[0]))+1)*-1.0/480.0, nil
}

func (l *LPS331AP) Active() error {
	id, err := l.Read(0x0f)
	if err != nil {
		return err
	}
	if id != 0xbb {
		return errors.New("Invalid device.")
	}

	if err != l.bus.Write(l.addr, 0x20, 0x90) {
		return err
	}

	l.active = true

	return nil
}

func (l *LPS331AP) Deactive() error {
	if !l.active {
		return nil
	}

	var err error
	if err != l.bus.Write(l.addr, 0x20, 0x0) {
		return err
	}

	l.active = false

	return nil
}
