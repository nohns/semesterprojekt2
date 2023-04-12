package domain

import (
	"context"
)

type domain struct {
	//uart uart
}

type uart interface {
	AwaitResponse(context.Context, string) (byte, error)
	Write([]byte) error
}

func New( /* uart uart */ ) *domain {
	return &domain{ /* uart: uart */ }
}

func (d domain) Register() (string, error) {
	return "", nil
}

func (d domain) GetLock() (bool, error) {

	//Convert "123" to byte
	/* ok := []byte("123")

	err := d.uart.Write(ok)
	if err != nil {
		return false, err
	}

	res, err := d.uart.AwaitResponse(context.Background(), "123")
	if err != nil {
		return false, err
	}
	//convert res to string and print
	fmt.Println(string(res)) */

	return false, nil
}

func (d domain) SetLock() (bool, error) {

	return false, nil
}
