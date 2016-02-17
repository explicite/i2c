package driver

import (
	"errors"
	"github.com/explicite/i2c"
)

type Driver struct {
	bus    *i2c.Bus
	addr   byte
	active bool
}

func (d *Driver) Load(addr byte, busNo byte) error {
	bus, err := i2c.NewBus(busNo)
	if err != nil {
		return err
	}
	d = &Driver{bus, addr, false}
	return nil
}

func (d *Driver) Read(cmd byte, length byte) (list []byte, err error) {
	return d.bus.Read(d.addr, cmd, length)
}

func (d *Driver) Write(cmd byte, list ...byte) (err error) {
	return d.bus.Write(d.addr, cmd, list...)
}

func (d *Driver) On() error {
	if d.active == true {
		return errors.New("device is on")
	}

	return nil
}

func (d *Driver) Off() error {
	if d.active != false {
		return errors.New("device is off")
	}
	return nil
}
