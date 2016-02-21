package bh1750

import (
	"github.com/explicite/i2c/driver"
	"time"
)

const (
	// ADDR ≦ 0.3VCC
	AddrL = 0x23

	// ADDR ≧ 0.7VCC
	AddrH = 0x5c

	// No active state.
	PowerDown = 0x00

	// Waiting for measurement command.
	PowerOn = 0x01

	// Reset Data register value. Reset command is not acceptable in Power Down mode.
	Reset = 0x07

	// Start measurement at 1lx resolution. Measurement Time is typically 120ms.
	ConHRes1lx = 0x10

	// Start measurement at 0.5lx resolution. Measurement Time is typically 120ms.
	ConHRes05lx = 0x11

	// Start measurement at 4lx resolution. Measurement Time is typically 16ms.
	ConLRes4lx = 0x13

	// Start measurement at 1lx resolution. Measurement Time is typically 120ms.
	// It is automatically set to Power Down mode after measurement.
	OtHRes1lx = 0x20

	// Start measurement at 0.5lx resolution. Measurement Time is typically 120ms.
	// It is automatically set to Power Down mode after measurement.
	OtHRes05lx = 0x21

	// Start measurement at 4lx resolution. Measurement Time is typically 16ms.
	// It is automatically set to Power Down mode after measurement.
	OtLRes4lx = 0x23

	// 20ms for safety time margine in measurement.
	Stm = 20 * time.Millisecond
)

// Map of timeouts for measurement type.
var timeout = map[byte]time.Duration{
	ConHRes1lx:  120*time.Millisecond + Stm,
	ConHRes05lx: 120*time.Millisecond + Stm,
	ConLRes4lx:  16*time.Millisecond + Stm,
	OtHRes1lx:   120*time.Millisecond + Stm,
	OtHRes05lx:  120*time.Millisecond + Stm,
	OtLRes4lx:   16*time.Millisecond + Stm,
}

type BH1750 struct{ driver.Driver }

func (b *BH1750) Init(addr byte, bus byte) error {
	return b.Load(addr, bus)
}

func (b *BH1750) Lux(mode byte) (float32, error) {
	b.Write(mode, 0x00)
	time.Sleep(timeout[mode])
	buf := make([]byte, 0x02)
	var err error
	buf, err = b.Read(mode, 0x02)

	if err != nil {
		return 0, err
	}

	return float32((int(buf[1]) + (int(buf[0]) >> 8))) / 1.2, nil
}

func (b *BH1750) Active() error {
	return b.On()
}

func (b *BH1750) Deactive() error {
	return b.Off()
}
