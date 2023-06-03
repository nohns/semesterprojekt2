package hw

import (
	"context"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const startPairBtnPin rpio.Pin = 21

type hw struct {
	prevState rpio.State
	c         hwConsumer
}

type hwConsumer interface {
	HandlePress()
}

func New(c hwConsumer) *hw {
	return &hw{
		prevState: rpio.Low,
		c:         c,
	}
}

func (h *hw) Listen() {
	rpio.Open()
	defer rpio.Close()

	// Set start pair pin to input with pullup
	startPairBtnPin.Mode(rpio.Input)
	startPairBtnPin.PullUp()

	log.Printf("listening for button presss")
	for {
		// Run the start pair handler if start pair button is pressed
		state := readPinDebounced(startPairBtnPin)
		if state == rpio.High && h.prevState == rpio.Low {
			h.c.HandlePress()
		}

		h.prevState = state
	}
}

func readPinDebounced(pin rpio.Pin) rpio.State {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start goroutine to read pin state updates concurrently
	states := make(chan rpio.State)
	go func() {
		prevState := pin.Read()
		for {
			// Stop reading if context is cancelled
			if ctx.Err() != nil {
				return
			}

			state := pin.Read()
			// Falling edge
			if state == rpio.Low && prevState == rpio.High {
				states <- state
			}
			// Rising edge
			if state == rpio.High && prevState == rpio.Low {
				states <- state
			}

			prevState = state
		}
	}()

	// Return only if state has been stable for 100ms
	currentState := pin.Read()
	for {
		select {
		case currentState = <-states:
			continue
		case <-time.After(100 * time.Millisecond):
			return currentState
		}
	}
}
