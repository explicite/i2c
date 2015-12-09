package bh1750

import (
	"github.com/explicite/i2c/driver"
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

	// 20ms for safety time margine in measurement.
	STM = 20 * time.Millisecond
)

// Map of timeouts for measurement type.
var timeout = map[byte]time.Duration{
	CON_H_RES_1LX:  120*time.Millisecond + STM,
	CON_H_RES_05LX: 120*time.Millisecond + STM,
	CON_L_RES_4LX:  16*time.Millisecond + STM,
	OT_H_RES_1LX:   120*time.Millisecond + STM,
	OT_H_RES_05LX:  120*time.Millisecond + STM,
	OT_L_RES_4LX:   16*time.Millisecond + STM,
}

type BH1750 struct {
	drv driver.Driver
}

func (b *BH1750) Init(addr byte, bus byte) error {
	var err error
	var drv *driver.Driver
	drv, err = driver.NewDriver(addr, bus)
	if err == nil {
		drv.Write(POWER_DOWN, 0x00)
		b.drv = *drv
	}

	return err
}

func (b *BH1750) Lux(mode byte) (float32, error) {
	b.drv.Write(mode, 0x00)
	time.Sleep(timeout[mode])
	buf := make([]byte, 0x02)
	var err error
	buf, err = b.drv.Read(mode, 0x02)

	if err != nil {
		return 0, err
	}

	return float32((uint8(buf[1]) + (uint8(buf[0] >> 8)))) / 1.2, nil
}

func (b *BH1750) Active() error {
	var err error
	if err = b.drv.On(); err != nil {
		return err
	}

	if err = b.drv.Write(POWER_ON, 0x00); err != nil {
		return err
	}

	return nil
}

func (b *BH1750) Deactive() error {
	var err error
	if err = b.drv.Off(); err != nil {
		return err
	}

	if err != b.drv.Write(POWER_DOWN, 0x00) {
		return err
	}

	return nil
}
