package uart

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"go.bug.st/serial"
)

// You need to manually check which port your arduino is connected to
// Could potentially be automated :TODO:
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
	port  serial.Port
	ch    chan []byte
	mappy map[string]string
	mu    *sync.RWMutex
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
	mappy := make(map[string]string)
	mu := &sync.RWMutex{}
	//Initialize the UART struct
	uart := &Uart{port: port, ch: dataCh, mappy: mappy, mu: mu}
	//Runs in a seperate goroutine and listens for data
	//Returns the data to the channel
	go uart.Listen()

	return uart
}

//Function that awaits the connect response based on ID

func (u *Uart) Write(b []byte) error {
	// Write some data to the UART device

	b = append(b, '\x00')
	fmt.Println("Writing", string(b))

	n, err := u.port.Write(b)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("Wrote %d bytes to the UART device.\r\n", n)
	return nil
}

func (u *Uart) Read() ([]byte, error) {
	result, err := bufio.NewReader(u.port).ReadBytes('\x00')
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println("Read", string(result))

	return result, nil
}

// Function that acts as a listerner for the UART should run in a seperate goroutine
func (u *Uart) Listen() {
	fmt.Println("Listening for data..")
	//Create go routine that listens for data

	for {
		data, err := u.Read()
		if err != nil {
			log.Println(err)
			continue
		}
		//Here it might potentially be better to store the request id and the response id in a map
		//Convert data to string
		dataString := string(data)
		//Extract the id from the data
		//Split by /
		splitter := strings.Split(dataString, "/")
		//index 0 is the id
		//Add the id to the map
		u.mappy[splitter[1]] = dataString

		u.ch <- data
	}

}

// method that listens in on the UART channel and looks for an id match with a value within the channel
func (u *Uart) AwaitResponse(ctx context.Context, id string) (byte, error) {
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
}

//The way I think I want this to work is that the server needs to listen in on the UART port
//So its gonna start out by sending a request to the uart and then the server should await a response which is blocking but it should be
//Cancelable if the context runs out
