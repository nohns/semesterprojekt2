package domain

import (
	"context"
	"log"
	"time"
)

// request constants
const openLock = 0b1011
const closeLock = 0b1010
const lockState = 0b1000

// Response constants
const ack = 0b1111
const nack = 0b1100
const locked = 0b1101
const unlocked = 0b1110

func (d domain) GetLock(ctx context.Context) (bool, error) {

	res, err := d.uart.AwaitResponse(ctx, lockState)
	if err != nil {
		return false, err
	}
	state, err := d.translator(ctx, res[0], lockState, false)
	if err != nil {
		return false, err
	}

	log.Println("Current state is: ", state)

	return state, nil
}

func (d domain) SetLock(ctx context.Context, state bool) (bool, error) {

	if state {
		res, err := d.uart.AwaitResponse(ctx, openLock)
		if err != nil {
			return false, err
		}
		state, err := d.translator(ctx, res[0], lockState, true)
		if err != nil {
			return false, err
		}

		log.Println("Current state is: ", state)
		return state, nil
	}
	if !state {
		res, err := d.uart.AwaitResponse(ctx, closeLock)
		if err != nil {
			return false, err
		}
		state, err := d.translator(ctx, res[0], lockState, false)
		if err != nil {
			return false, err
		}

		log.Println("Current state is: ", state)
		return state, nil
	}

	return false, nil
}

// Function that compares the result of the result with the command constants and returns the correct state
func (d domain) translator(ctx context.Context, b byte, cmd int, state bool) (bool, error) {

	if ctx.Err() != nil {
		log.Println("Context has been canceled")
		// Context has been canceled, return error or appropriate value
		return false, ctx.Err()
	}

	if b == ack {
		//Here we know the response is the same as the one we set
		return state, nil
	}

	if b == locked {
		//Here we know the door is locked so we return true
		return true, nil
	}

	if b == unlocked {
		//Here we know the door is unlocked so we return false
		return false, nil

	}

	if b == nack {
		//Retry sending the cmd
		res, err := d.uart.AwaitResponse(ctx, cmd)
		if err != nil {
			return false, err
		}
		//wait 300ms and then retry sending the cmd
		time.Sleep(300 * time.Millisecond)
		return d.translator(ctx, res[0], cmd, state)
	}

	//We didn't recieve any valid command at all so we retry sending the cmd
	res, err := d.uart.AwaitResponse(ctx, cmd)
	if err != nil {
		return false, err
	}
	//wait 300ms and then retry sending the cmd
	time.Sleep(300 * time.Millisecond)
	return d.translator(ctx, res[0], cmd, state)

}
