package si7021

import (
	"github.com/explicite/i2c/driver"
)

const (
	// Measure Relative Humidity, Hold Master Mode.
	RhHm = 0xE5

	// Measure Relative Humidity, No Hold Master Mode.
	Rh = 0xF5

	// Measure Temperature, Hold Master Mode.
	TmpHm = 0xE3

	// Measure Temperature, No Hold Master Mode.
	Tmp = 0xF3

	// Read Temperature Value from Previous RH Measurement.
	ReadTmpPrev = 0xE0

	// Reset.
	Reset = 0xFE

	// Write RH/T User Register 1
	WriteRhtUr1 = 0xE6

	// Read RH/T User Register 1
	ReadRhtUr1 = 0xE7

	// Write Heater Control Register
	WriteHcr = 0x51

	// Read Heater Control Register
	ReadHcr = 0x11

	// Read Electronic ID 1st Byte
	ReadEid1p1 = 0xFA
	ReadEid1p2 = 0x0F

	// Read Electronic ID 2nd Byte
	ReadEid2p1 = 0xFC
	ReadEid2p2 = 0xC9

	// Read Firmware Revision
	ReadFr1 = 0x84
	ReadFr2 = 0xB8
)

type SI7021 struct{ driver.Driver }

func (s *SI7021) Init(addr byte, bus byte) error {
	return s.Load(addr, bus)
}

func (s *SI7021) RelativeHumidity(hm bool) (float64, error) {
	//TODO
	if hm == true {
		s.Write(RhHm, 0x01)
	} else {
		s.Write(Rh, 0x01)
	}
	return float64(1), nil
}

func (s *SI7021) Temperature(hm bool) (float64, error) {
	//TODO
	if hm == true {
		s.Write(TmpHm, 0x01)
	} else {
		s.Write(Tmp, 0x01)
	}
	return float64(1), nil
}

func (s *SI7021) ESN() (string, error) {
	//TODO
	return "todo", nil
}

func (s *SI7021) Rev() (string, error) {
	//TODO
	return "todo", nil
}

func (s *SI7021) Active() error {
	return s.On()
}

func (s *SI7021) Deactive() error {
	return s.Off()
}
