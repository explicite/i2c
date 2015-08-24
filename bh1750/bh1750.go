package bh1750

import (
	"time"

	"github.com/explicite/i2c"
)

const (
	WRITE_REG = 0x46

	READ_REG = 0x47

	//No active state.
	POWER_DOWN = 0x00

	//Waiting for measurement command.
	POWER_ON = 0x01

	//Reset Data register value. Reset command is not acceptable in Power Down mode.
	RESET = 0x07

	//Start measurement at 1lx resolution. Measurement Time is typically 120ms.
	CON_H_RES_1LX = 0x10

	//Start measurement at 0.5lx resolution. Measurement Time is typically 120ms.
	CON_H_RES_05LX = 0x11

	//Start measurement at 4lx resolution. Measurement Time is typically 16ms.
	CON_L_RES_4LX = 0x13

	//Start measurement at 1lx resolution. Measurement Time is typically 120ms. It is automatically set to Power Down mode after measurement.
	OT_H_RES_1LX = 0x20

	//Start measurement at 0.5lx resolution. Measurement Time is typically 120ms. It is automatically set to Power Down mode after measurement.
	OT_H_RES_05LX = 0x21

	//Start measurement at 4lx resolution. Measurement Time is typically 16ms. It is automatically set to Power Down mode after measurement.
	OT_L_RES_4LX = 0x23
)

type BH1750 struct {
	bus    *i2c.Bus
	active bool
}

func (b *BH1750) Init(addr byte, bus byte) error {
	var err error
	b.bus, err = i2c.NewBus(addr, bus)
	b.bus.Write(WRITE_REG, POWER_DOWN)
	return err
}

func (b *BH1750) Lux(mode byte) (float32, error) {
	b.bus.Write(mode, 0x00)
	time.Sleep(120 * time.Millisecond)
	buf := make([]byte, 2)
	var err error
	buf, err = b.bus.Read(WRITE_REG, 2)

	if err != nil {
		return 0, err
	}

	return float32((int16(buf[1]) + (256 * int16(buf[0])))) / 1.2, nil
}

func (b *BH1750) Active() error {
	var err error
	if err != b.bus.Write(WRITE_REG, POWER_ON) {
		return err
	}

	b.active = true

	return nil
}

func (b *BH1750) Deactive() error {
	if !b.active {
		return nil
	}

	var err error
	if err != b.bus.Write(WRITE_REG, POWER_DOWN) {
		return err
	}

	b.active = false

	return nil
}
