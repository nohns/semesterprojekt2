package uart

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

const USB = "/dev/tty.usbmodem14501"

func Uart() {
	// Open the serial port
	port, err := serial.Open("/dev/tty.usbmodem14501", &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		StopBits: serial.OneStopBit,
		Parity:   serial.NoParity,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Write some data to the UART device
	data := []byte("b")
	n, err := port.Write(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes to the UART device.\n", n)

	// Read some data from the UART device
	buffer := make([]byte, 100)
	n, err = port.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes from the UART device: %s\n", n, string(buffer[:n]))

	// Wait for a second
	//Loop forever
	for {
		data := []byte("b")
		n, err := port.Write(data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Wrote %d bytes to the UART device.\n", n)
		buffer := make([]byte, 100)
		n, err = port.Read(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Read %d bytes from the UART device: %s\n", n, string(buffer[:n]))

		time.Sleep(1 * time.Second)
	}
}

//I think I want to have a seperate goroutine running as reader
// and I should have a goroutine running as writer

//Real abstraction layer for the UART

type uart interface {
	read() (byte, error)
	write(byte) (byte, error)
}

type uartImpl struct {
	port serial.Port
}

func New() *uartImpl {
	return &uartImpl{}
}

func (u *uartImpl) Read() (byte, error) {
	// Read some data from the UART device
	buffer := make([]byte, 100)
	n, err := u.port.Read(buffer)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Printf("Read %d bytes from the UART device: %s\n", n, string(buffer[:n]))
	return buffer[0], nil
}

func (u *uartImpl) Write(b byte) (byte, error) {
	// Write some data to the UART device
	data := []byte{b}
	n, err := u.port.Write(data)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Printf("Wrote %d bytes to the UART device.\n", n)
	return data[0], nil
}
