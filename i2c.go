package i2c

import (
	"fmt"
	"os"
	"sync"
	"syscall"
	"unsafe"
)

// as defined in /usr/include/linux/i2c-dev.h
const (
	I2C_SLAVE = 0x0703
	I2C_SMBUS = 0x0720
)

// as defined in /usr/include/linux/i2c.h
const (
	I2C_SMBUS_WRITE          = 0
	I2C_SMBUS_READ           = 1
	I2C_SMBUS_I2C_BLOCK_DATA = 8
	I2C_SMBUS_BLOCK_MAX      = 32
)

var busMap map[byte]*Bus
var busMapLock sync.Mutex

func init() {
	busMap = make(map[byte]*Bus)
}

// as defined in /usr/include/linux/i2c-dev.h
type i2c_smbus_ioctl_data struct {
	readWrite byte
	command   byte
	size      uint32
	data      uintptr
}

type Bus struct {
	// i2c-dev file pointer
	file *os.File

	// simple bus access lock to ensure address
	// set and data writes occur atomically
	lock sync.Mutex
}

// Returns an instance to an I2CBus. If we already have an I2CBus
// created for the requested bus number, just return that, otherwise
// set up a new one and open up its associated i2c-dev file
func NewBus(bus byte) (i2cbus *Bus, err error) {
	busMapLock.Lock()
	defer busMapLock.Unlock()

	if i2cbus = busMap[bus]; i2cbus == nil {
		i2cbus = new(Bus)
		file, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%v", bus), os.O_RDWR, os.ModeDevice)

		if err != nil {
			return i2cbus, err
		}

		i2cbus.file = file

		busMap[bus] = i2cbus
		return i2cbus, nil
	}

	return
}

func (i2cbus *Bus) SetAddress(addr byte) (err error) {
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, i2cbus.file.Fd(), I2C_SLAVE, uintptr(addr)); errno != 0 {
		err = syscall.Errno(errno)
		return
	}

	return
}

func (i2cbus *Bus) Read(addr byte, cmd byte, length byte) (list []byte, err error) {
	i2cbus.lock.Lock()
	defer i2cbus.lock.Unlock()

	if err := i2cbus.SetAddress(addr); err != nil {
		panic(err)
	}

	data := make([]byte, length+1)
	data[0] = length

	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		i2cbus.file.Fd(), I2C_SMBUS, uintptr(unsafe.Pointer(&i2c_smbus_ioctl_data{
			readWrite: I2C_SMBUS_READ,
			command:   cmd,
			size:      I2C_SMBUS_I2C_BLOCK_DATA,
			data:      uintptr(unsafe.Pointer(&data[0]))}))); errno != 0 {
		err = syscall.Errno(errno)
	}

	list = make([]byte, length)
	copy(list, data[1:])

	return
}

func (i2cbus *Bus) Write(addr byte, cmd byte, list ...byte) (err error) {
	i2cbus.lock.Lock()
	defer i2cbus.lock.Unlock()

	if err := i2cbus.SetAddress(addr); err != nil {
		panic(err)
	}

	data := make([]byte, len(list)+1)
	data[0] = byte(len(list))
	copy(data[1:], list)

	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		i2cbus.file.Fd(), I2C_SMBUS, uintptr(unsafe.Pointer(&i2c_smbus_ioctl_data{
			readWrite: I2C_SMBUS_WRITE,
			command:   cmd,
			size:      I2C_SMBUS_I2C_BLOCK_DATA,
			data:      uintptr(unsafe.Pointer(&data[0]))}))); errno != 0 {
		err = syscall.Errno(errno)
	}

	return
}
