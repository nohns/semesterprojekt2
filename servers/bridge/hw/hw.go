package hw

import (
	"sync"

	"github.com/stianeikeland/go-rpio/v4"
)

const startPairBtnPin rpio.Pin = 5

var defaultHW = &hw{}

type hw struct {
	startPairHandler func()
	mu               sync.Mutex
}

func (h *hw) listen() {
	rpio.Open()
	defer rpio.Close()

	rpio.PinMode(startPairBtnPin, rpio.Input)

	for {
		// Run the start pair handler if start pair button is pressed
		h.mu.Lock()
		if rpio.ReadPin(startPairBtnPin) == rpio.Low && h.startPairHandler != nil {
			h.startPairHandler()
		}
		h.mu.Unlock()
	}

}

func (hw *hw) handlePairStartPress(h func()) {
	hw.mu.Lock()
	defer hw.mu.Unlock()

	hw.startPairHandler = h
}

// Listen for hardware changes, like pair button presses
func Listen() {
	defaultHW.listen()
}

// HandlePairStartPress registers a handler for when the pair button is pressed
func HandlePairStartPress(h func()) {
	defaultHW.handlePairStartPress(h)
}
