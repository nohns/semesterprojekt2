package uart

import (
	"bufio"
	"fmt"
	"log"
	"sync"

	"go.bug.st/serial"
)

// You need to manually check which port your arduino is connected to
// Could potentially be automated :TODO:
// const USB = "/dev/ttyACM0"
const USB = "/dev/tty.usbmodem14501"

//TODO: All of this code needs to be fixed up
//I feel fairly certain that a race condition is happening here

type Uart struct {
	port serial.Port

	mu sync.Mutex
}

func New() *Uart {
	port, err := serial.Open(USB, &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		StopBits: serial.OneStopBit,
		Parity:   serial.NoParity,
	})
	if err != nil {
		//log.Fatal("Failed to open UART connection", err)
	}

	//Initialize the UART struct
	uart := &Uart{port: port, mu: sync.Mutex{}}

	return uart
}

// Function that writes b bytes to the UART device and adds a null terminator
func (u *Uart) Write(b []byte) error {

	//Append null terminator to the end of the byte array
	b = append(b, '\x00')
	fmt.Printf("Wrote %b to the UART device.\r\n", b)

	//Write the bytes to the UART device
	n, err := u.port.Write(b)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Wrote %d bytes\r\n", n)

	return nil
}

// Function that reads from the UART device until it encounters a null terminator
func (u *Uart) Read() ([]byte, error) {

	result, err := bufio.NewReader(u.port).ReadBytes('\x00')
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("Response received %b\r\n", result)
	return result, nil
}

func (u *Uart) AwaitResponse(cmd int) ([]byte, error) {
	//Lock the mutex to ensure that no other goroutine is writing to the UART device
	u.mu.Lock()
	defer u.mu.Unlock()

	//Convert command to a byte
	b := []byte{byte(cmd)}
	//Call the write function
	err := u.Write(b)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	//await response
	res, err := u.Read()
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return res, nil
}
