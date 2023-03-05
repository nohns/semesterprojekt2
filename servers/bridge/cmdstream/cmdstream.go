package cmdstream

import (
	"context"
	"log"

	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
)

type stream struct {
	c bridgepb.CmdServiceClient
	d *distributor
}

func New(c bridgepb.CmdServiceClient, d *distributor) *stream {
	return &stream{
		c: c,
		d: d,
	}
}

// Starts new stream of commands. This function blocks until the stream is closed.
func (s *stream) Listen(ctx context.Context) error {
	sc, err := s.c.StreamCommands(ctx, &bridgepb.StreamCommandsRequest{})
	if err != nil {
		return err
	}

	// Receive in this one and block until a potential error occurs
	return s.recv(sc.Context(), sc)
}

// Receives commands from the cloud and distributes them to the correct services
func (s *stream) recv(ctx context.Context, sc bridgepb.CmdService_StreamCommandsClient) error {
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
