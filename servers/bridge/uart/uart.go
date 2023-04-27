package uart

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"sync"

	"go.bug.st/serial"
)

// You need to manually check which port your arduino is connected to
// Could potentially be automated :TODO:
const USB = "/dev/tty.usbmodem14501"

//TODO: All of this code needs to be fixed up
//I feel fairly certain that a race condition is happening here

type Uart struct {
	port serial.Port
	ch   chan []byte
	mu   sync.Mutex
}

func New() *Uart {
	port, err := serial.Open(USB, &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		StopBits: serial.OneStopBit,
		Parity:   serial.NoParity,
	})
	if err != nil {
		log.Fatal("Failed to open UART connection", err)
	}

	dataCh := make(chan []byte, 100)

	//Initialize the UART struct
	uart := &Uart{port: port, ch: dataCh, mu: sync.Mutex{}}
	//Runs in a seperate goroutine and listens for data
	//Returns the data to the channel
	//go uart.Listen()

	return uart
}

// Function that writes b bytes to the UART device and adds a null terminator
func (u *Uart) write(b []byte) error {

	//Append null terminator to the end of the byte array
	b = append(b, '\x00')
	fmt.Printf("Wrote %b to the UART device.\r\n", b)

	//Write the bytes to the UART device
	n, err := u.port.Write(b)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Wrote %d bytes to the UART device.\r\n", n)

	return nil
}

// Function that reads from the UART device until it encounters a null terminator
func (u *Uart) read() ([]byte, error) {
	result, err := bufio.NewReader(u.port).ReadBytes('\x00')
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (u *Uart) AwaitResponse(ctx context.Context, cmd int) ([]byte, error) {
	//Lock the mutex to ensure that no other goroutine is writing to the UART device
	u.mu.Lock()
	defer u.mu.Unlock()

	//Convert command to a byte
	b := []byte{byte(cmd)}
	//Call the write function
	err := u.write(b)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	//await response
	res, err := u.read()
	if err != nil {
		log.Println(err)

		return nil, err
	}
	fmt.Printf("Response received %b\r\n", res)
	return res, nil
}

// Function that acts as a listerner for the UART should run in a seperate goroutine
/* func (u *Uart) Listen() {
	fmt.Println("Listening for data..")
	//Create go routine that listens for data

	for {
		data, err := u.Read()
		if err != nil {
			log.Println(err)
			continue
		}
		for _, n := range data {
			fmt.Printf("%08b ", n) // prints 00000000 11111101
		}

		u.ch <- data
	}
} */

//The way I think I want this to work is that the server needs to listen in on the UART port
//So its gonna start out by sending a request to the uart and then the server should await a response which is blocking but it should be
//Cancelable if the context runs out

// method that listens in on the UART channel and looks for an id match with a value within the channel
/* func (u *Uart) AwaitResponse(ctx context.Context, id string) (byte, error) {
	fmt.Println("Awaiting response")
	for {
		select {

		case data := <-u.ch:

			//Check if data contains the correct id
			if string(data) == id {
				return data[1], nil
			}

		default:
			return 0, errors.New("no response received")
		}
	}
} */
