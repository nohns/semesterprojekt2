package uart

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

func main() {
    // Open the serial port
    port, err := serial.Open("/dev/ttyUSB0", &serial.Mode{
        BaudRate: 115200,
        DataBits: 8,
        StopBits: serial.OneStopBit,
        Parity:   serial.NoParity,
    })
    if err != nil {
        panic(err)
    }
    defer port.Close()

    // Write some data to the UART device
    data := []byte("Hello, UART!")
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
    time.Sleep(time.Second)
}