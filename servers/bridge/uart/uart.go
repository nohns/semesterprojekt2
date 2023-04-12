package uart

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.bug.st/serial"
)

const USB = "/dev/tty.usbmodem14501"

/* func Uart() {
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
} */

//I think I want to have a seperate goroutine running as reader
// and I should have a goroutine running as writer

//Real abstraction layer for the UART

//TODO: All of this code needs to be fixed up
//I feel fairly certain that a race condition is happening here

type Uart struct {
	port serial.Port
	ch   chan []byte
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
	uart := &Uart{port: port, ch: dataCh}

	//Runs in a seperate goroutine and listens for data
	//Returns the data to the channel
	uart.Listen()

	return uart
}

func (u *Uart) Read() (byte, error) {
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

//Function that awaits the connect response based on ID

func (u *Uart) Write(b []byte) error {
	// Write some data to the UART device

	n, err := u.port.Write(b)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Wrote %d bytes to the UART device.\n", n)
	return nil
}

// Function that acts as a listerner for the UART should run in a seperate goroutine
func (u *Uart) Listen() {
	fmt.Println("Listening for data..")
	//Create go routine that listens for data
	go func() {
		for {
			data, err := u.Read()
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("Data received", data)
			u.ch <- []byte{data}
		}
	}()
}

// method that listens in on the UART channel and looks for an id match with a value within the channel
func (u *Uart) AwaitResponse(ctx context.Context, id string) (byte, error) {
	fmt.Println("Awaiting response")
	for {
		select {
		case data := <-u.ch:
			fmt.Println("Data received", data)
			if len(data) < 1 {
				return 0, errors.New("invalid data received")
			}
			if data[0] == id[0] {
				return data[0], nil
			}
		default:
			return 0, errors.New("no response received")
		}
	}
}

//The way I think I want this to work is that the server needs to listen in on the UART port
//So its gonna start out by sending a request to the uart and then the server should await a response which is blocking but it should be
//Cancelable if the context runs out
