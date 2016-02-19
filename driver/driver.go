package driver

import (
	"errors"
	"github.com/explicite/i2c"
)

// Driver structure for handling i2c standard functions.
type Driver struct {
	bus    *i2c.Bus
	addr   byte
	active bool
}

// Load(addr, busNo) initialize i2c bus and save device address.
func (d *Driver) Load(addr byte, busNo byte) error {
	bus, err := i2c.NewBus(busNo)
	if err != nil {
		return err
	}
	*d = Driver{bus, addr, false}
	return nil
}

// Read(cmd, length) read data[length] from device register.
func (d *Driver) Read(cmd byte, length byte) (list []byte, err error) {
	return d.bus.Read(d.addr, cmd, length)
}

// Write(cmd, list) write data to device register.
func (d *Driver) Write(cmd byte, list ...byte) (err error) {
	return d.bus.Write(d.addr, cmd, list...)
}

// On() device - TODO should be implemented locking
func (d *Driver) On() error {
	if d.active == true {
		return errors.New("device is on")
	}

	return nil
}

// Off() device - TODO should be implemented locking
func (d *Driver) Off() error {
	if d.active != false {
		return errors.New("device is off")
	}
	return nil
}
