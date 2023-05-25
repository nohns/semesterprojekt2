package hw

import (
	"github.com/stianeikeland/go-rpio/v4"
)

const startPairBtnPin rpio.Pin = 5

type hw struct {
	c hwConsumer
}

type hwConsumer interface {
	HandlePress()
}

func New(c hwConsumer) *hw {
	return &hw{
		c: c,
	}
}

func (h *hw) Listen() {
	rpio.Open()
	defer rpio.Close()

	rpio.PinMode(startPairBtnPin, rpio.Input)

	for {
		// Run the start pair handler if start pair button is pressed
		if rpio.ReadPin(startPairBtnPin) == rpio.Low {
			h.c.HandlePress()
		}
	}
}
