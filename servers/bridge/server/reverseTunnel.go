package server

import (
	"context"
	"log"

	"github.com/jhump/grpctunnel"
	"github.com/jhump/grpctunnel/tunnelpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func ServeTunnel(s *grpc.Server, server *Server) {
	conn, err := grpc.Dial(
		//"127.0.0.1:8000",
		"dns:///0.0.0.0"+":8500",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	tunnelStub := tunnelpb.NewTunnelServiceClient(conn)
	channelServer := grpctunnel.NewReverseTunnelServer(tunnelStub)

	//Register the server
	pairingv1.RegisterPairingServiceServer(s, server)
	lockv1.RegisterLockServiceServer(s, server)

	//Create channel to listen in to

	log.Println("Registering tunnel server")
	if _, err := channelServer.Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

/*
func NewTunnelServer() {

	handler := grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{},
	)
	svr := grpc.NewServer()
	tunnelpb.RegisterTunnelServiceServer(svr, handler.Service())

	// TODO: Configure services for forward tunnels.
	// TODO: Inject handler into code that will use reverse tunnels.

	// Start the gRPC server.
	l, err := net.Listen("tcp", "0.0.0.0:8500")
	if err != nil {
		log.Fatal(err)
	}
	if err := svr.Serve(l); err != nil {
		log.Fatal(err)
	}
} */
