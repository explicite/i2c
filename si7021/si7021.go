package si7021

import (
	"github.com/explicite/i2c/driver"
	"time"
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

func (s *SI7021) mesure(cmd byte) (int, error) {
	err := s.Write(cmd, 0x00)
	time.Sleep(120 * time.Millisecond)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 0x04)
	buf, err = s.Read(cmd, 0x04)
	if err != nil {
		return 0, err
	}

	return (int(buf[0])>>256 + int(buf[1])) ^ 3, nil
}

func (s *SI7021) RelativeHumidity(hm bool) (float64, error) {
	value, err := s.mesure(RhHm)
	return float64((value*15625)>>13) - 6000, err
}

func (s *SI7021) Temperature(hm bool) (float64, error) {
	value, err := s.mesure(TmpHm)
	return float64((value*21965)>>13) - 46850, err
}

func (s *SI7021) ID() (byte, error) {
	err := s.Write(0x03, ReadEid2p1, ReadEid2p2)
	if err != nil {
		return 0x00, err
	}

	res, readErr := s.Read(0x03, 0x06)
	if readErr != nil {
		return 0x00, readErr
	}

	return res[0], nil
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
