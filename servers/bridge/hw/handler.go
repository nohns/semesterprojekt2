package hw

import (
	"context"
	"log"
	"time"
)

// Define a timeout duration for button press, until we should stop listening for phone handshake
const handshakeTimeoutDuration = 30 * time.Second

type handshakeInitiator interface {
	BeginHandshake(ctx context.Context) error
}

type buttonHandler struct {
	handshaker handshakeInitiator
}

func NewButtonHandler(handshaker handshakeInitiator) *buttonHandler {
	return &buttonHandler{
		handshaker: handshaker,
	}
}

func (h *buttonHandler) HandlePress() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), handshakeTimeoutDuration)
	defer cancel()

	// Start handshake
	if err := h.handshaker.BeginHandshake(ctx); err != nil {
		log.Printf("failed to begin handshake: %v", err)
		return
	}
}
