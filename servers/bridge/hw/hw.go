package hw

import (
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
		prevState: rpio.High,
		c:         c,
	}
}

func (h *hw) Listen() {
	rpio.Open()
	defer rpio.Close()

	// Set start pair pin to input with pullup
	startPairBtnPin.Mode(rpio.Input)
	startPairBtnPin.PullUp()

	for {
		// Run the start pair handler if start pair button is pressed
		state := startPairBtnPin.Read()
		if state == rpio.Low && h.prevState == rpio.High {
			h.c.HandlePress()
		}

		h.prevState = state
	}
}
