package si7021

import (
	"github.com/explicite/i2c/driver"
)

const (
	// RhHm code Measure Relative Humidity, Hold Master Mode.
	RhHm = 0xE5

	// Rh code Measure Relative Humidity, No Hold Master Mode.
	Rh = 0xF5

	// TmpHm code Measure Temperature, Hold Master Mode.
	TmpHm = 0xE3

	// Tmp code Measure Temperature, No Hold Master Mode.
	Tmp = 0xF3

	// ReadTmpPrev code Read Temperature Value from Previous RH Measurement.
	ReadTmpPrev = 0xE0

	// Reset code.
	Reset = 0xFE

	// WriteRhtUr1 code Write RH/T User Register 1
	WriteRhtUr1 = 0xE6

	// ReadRhtUr1 code Read RH/T User Register 1
	ReadRhtUr1 = 0xE7

	// WriteHcr code Write Heater Control Register
	WriteHcr = 0x51

	// ReadHcr code Read Heater Control Register
	ReadHcr = 0x11

	// ReadEid1p1 code Read Electronic ID 1st Byte part 1
	ReadEid1p1 = 0xFA
	// ReadEid1p2 code Read Electronic ID 1st Byte part 2
	ReadEid1p2 = 0x0F

	// ReadEid2p1 code Read Electronic ID 2nd Byte part 1
	ReadEid2p1 = 0xFC
	// ReadEid2p2 code Read Electronic ID 2nd Byte part 2
	ReadEid2p2 = 0xC9

	// ReadFr1 code Read Firmware Revision part 1
	ReadFr1 = 0x84
	// ReadFr2 code Read Firmware Revision part 2
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
