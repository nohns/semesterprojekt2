package domain

import (
	"fmt"
	"time"
)

func (d domain) GetLock() (bool, error) {

	//Convert "123" to byte
	//ok := []byte("GET/123/true\x00")
	ok := []byte("GET/123/true")

	err := d.uart.Write(ok)
	if err != nil {
		return false, err
	}
	time.Sleep(100 * time.Millisecond)

	res, err := d.uart.Read()
	if err != nil {
		return false, err
	}

	/* res, err := d.uart.AwaitResponse(context.Background(), "123")
	if err != nil {
		return false, err
	} */
	//convert res to string and print
	fmt.Println(string(res))

	return false, nil
}

func (d domain) SetLock() (bool, error) {

	return false, nil
}
