package querystream

import (
	"context"

	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
)

type stream struct {
	c bridgepb.QueryServiceClient
}

func New(c bridgepb.QueryServiceClient) *stream {
	return &stream{
		c: c,
	}
}

// Starts new stream of commands. This function blocks until the stream is closed.
func (s *stream) Listen(ctx context.Context) error {
	sqc, err := s.c.StreamQueries(ctx)
	if err != nil {
		return err
	}

	// Receive in this one and block until a potential error occurs
	return s.recv(sqc.Context(), sqc)
}

// Receives commands from the cloud and distributes them to the correct services
func (s *stream) recv(ctx context.Context, sqc bridgepb.QueryService_StreamQueriesClient) error {
	for {
		_, err := sqc.Recv()
		if err != nil {
			return err
		}

		/*if err := s.d.distribute(ctx, streamcmd); err != nil {
			log.Printf("error distributing command %s: %v", streamcmd.Type, err)
		}*/
	}
}
