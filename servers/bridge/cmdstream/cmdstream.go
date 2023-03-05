package cmdstream

import (
	"context"
	"log"

	"github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
)

type stream struct {
	c       bridge.BridgeServiceClient
	d       *distributor
	answers chan bridge.StreamAnswer
}

func New(c bridge.BridgeServiceClient, d *distributor) *stream {
	return &stream{
		c: c,
		d: d,
	}
}

// Starts new stream of commands. This function blocks until the stream is closed.
func (s *stream) Listen(ctx context.Context) error {
	sc, err := s.c.StreamActions(ctx)
	if err != nil {
		return err
	}

	// Send answers in a separate goroutine, and receive in this one.
	go s.send(sc.Context(), sc)
	return s.recv(sc.Context(), sc)
}

// Receives answers from the bridge and sends them to the cloud
func (s *stream) send(ctx context.Context, sc bridge.BridgeService_StreamActionsClient) error {
	for answer := range s.answers {
		if err := sc.Send(&answer); err != nil {
			return err
		}
	}
	return nil
}

// Receives commands from the cloud and distributes them to the correct services
func (s *stream) recv(ctx context.Context, sc bridge.BridgeService_StreamActionsClient) error {
	for {
		streamcmd, err := sc.Recv()
		if err != nil {
			return err
		}

		if err := s.d.distribute(ctx, streamcmd); err != nil {
			log.Printf("error distributing command %s: %v", streamcmd.Type, err)
		}
	}
}
