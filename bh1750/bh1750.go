package bh1750

import (
	"github.com/explicite/i2c"
	"time"
)

const (
	// ADDR ≦ 0.3VCC
	ADDR_L = 0x23

	// ADDR ≧ 0.7VCC
	ADDR_H = 0x5c

	// No active state.
	POWER_DOWN = 0x00

	// Waiting for measurement command.
	POWER_ON = 0x01

	// Reset Data register value. Reset command is not acceptable in Power Down mode.
	RESET = 0x07

	// Start measurement at 1lx resolution. Measurement Time is typically 120ms.
	CON_H_RES_1LX = 0x10

	// Start measurement at 0.5lx resolution. Measurement Time is typically 120ms.
	CON_H_RES_05LX = 0x11

	// Start measurement at 4lx resolution. Measurement Time is typically 16ms.
	CON_L_RES_4LX = 0x13

	// Start measurement at 1lx resolution. Measurement Time is typically 120ms.
	// It is automatically set to Power Down mode after measurement.
	OT_H_RES_1LX = 0x20

	// Start measurement at 0.5lx resolution. Measurement Time is typically 120ms.
	// It is automatically set to Power Down mode after measurement.
	OT_H_RES_05LX = 0x21

	// Start measurement at 4lx resolution. Measurement Time is typically 16ms.
	// It is automatically set to Power Down mode after measurement.
	OT_L_RES_4LX = 0x23
)

var timeout = map[byte]time.Duration{
	CON_H_RES_1LX:  120 * time.Millisecond,
	CON_H_RES_05LX: 120 * time.Millisecond,
	CON_L_RES_4LX:  16 * time.Millisecond,
	OT_H_RES_1LX:   120 * time.Millisecond,
	OT_H_RES_05LX:  120 * time.Millisecond,
	OT_L_RES_4LX:   16 * time.Millisecond,
}

type BH1750 struct {
	bus    *i2c.Bus
	addr   byte
	active bool
}

func (b *BH1750) Init(addr byte, bus byte) error {
	var err error
	b.bus, err = i2c.NewBus(bus)
	b.addr = addr
	b.bus.Write(addr, POWER_DOWN, 0x00)

	return err
}

func (b *BH1750) Lux(mode byte) (float32, error) {
	b.bus.Write(b.addr, mode, 0x00)
	time.Sleep(timeout[mode])
	buf := make([]byte, 0x02)
	var err error
	buf, err = b.bus.Read(b.addr, mode, 0x02)

	if err != nil {
		return 0, err
	}

	return float32((uint8(buf[1]) + (uint8(buf[0] >> 8)))) / 1.2, nil
}

func (b *BH1750) Active() error {
	var err error
	if err != b.bus.Write(b.addr, POWER_ON, 0x00) {
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
	if err != b.bus.Write(b.addr, POWER_DOWN, 0x00) {
		return err
	}

	b.active = false

	return nil
}
