package si7021

import (
	"github.com/explicite/i2c/driver"
)

const (
	// Measure Relative Humidity, Hold Master Mode.
	READ_MRH_HM = 0xE5

	// Measure Relative Humidity, No Hold Master Mode.
	READ_MRH = 0xF5

	// Measure Temperature, Hold Master Mode.
	READ_TMP_HM = 0xE3

	// Measure Temperature, No Hold Master Mode.
	READ_TMP = 0xF3

	// Read Temperature Value from Previous RH Measurement.
	READ_TMP_PREV = 0xE0

	// Reset.
	RESET = 0xFE

	// Write RH/T User Register 1
	WRITE_RHT_UR1 = 0xE6

	// Read RH/T User Register 1
	READ_RHT_UR1 = 0xE7

	// Write Heater Control Register
	WRITE_HCR = 0x51

	// Read Heater Control Register
	READ_HCR = 0x11

	// Read Electronic ID 1st Byte
	READ_EID_1_1 = 0xFA
	READ_EID_1_2 = 0x0F

	// Read Electronic ID 2nd Byte
	READ_EID_2_1 = 0xFC
	READ_EID_2_2 = 0xC9

	// Read Firmware Revision
	READ_FR_1 = 0x84
	READ_FR_2 = 0xB8
)

type SI7021 struct {
	drv driver.Driver
}
