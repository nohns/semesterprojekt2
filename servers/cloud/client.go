package cloud

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/jhump/grpctunnel"
	"github.com/jhump/grpctunnel/tunnelpb"
	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func NewLockClient() *lockv1.LockServiceClient /* , *grpc.ClientConn */ {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("BRIDGE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := lockv1.NewLockServiceClient(conn)
	return &client //, conn
}

func NewPairingClient() *pairingv1.PairingServiceClient /* , *grpc.ClientConn */ {

	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("BRIDGE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := pairingv1.NewPairingServiceClient(conn)
	return &client //, conn
}

func NewTunnel(conn *grpc.ClientConn) error {
	log.Println("Tunnel server listening on: " + os.Getenv("TUNNEL"))
	tunnelStub := tunnelpb.NewTunnelServiceClient(conn)
	channelServer := grpctunnel.NewReverseTunnelServer(tunnelStub)
	lockv1.RegisterLockServiceServer(channelServer, lockv1.UnimplementedLockServiceServer{})

	if _, err := channelServer.Serve(context.Background()); err != nil {
		log.Printf("failed to start tunnel server: %v", err)
		return err
	}
	return nil
}

func NewTunnelServer() {
	handler := grpctunnel.NewTunnelServiceHandler(
		grpctunnel.TunnelServiceHandlerOptions{},
	)
	svr := grpc.NewServer()
	tunnelpb.RegisterTunnelServiceServer(svr, handler.Service())

	// TODO: Configure services for forward tunnels.
	// TODO: Inject handler into code that will use reverse tunnels.

	// Start the gRPC server.
	l, err := net.Listen("tcp", "0.0.0.0:7899")
	if err != nil {
		log.Fatal(err)
	}
	if err := svr.Serve(l); err != nil {
		log.Fatal(err)
	}
}
