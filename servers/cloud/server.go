package cloud

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/nohns/servers/pkg/middleware"

	lockv1 "github.com/nohns/proto/lock/v1"

	pairingv1 "github.com/nohns/proto/pairing/v1"
)

//Server responbilbe for communication with the react native phone app

// This struct should take in
type server struct {
	lockv1.UnimplementedLockServiceServer
	pairingv1.UnimplementedPairingServiceServer

	lockClient    lockv1.LockServiceClient
	pairingClient pairingv1.PairingServiceClient
}

func newServer(lockClient lockv1.LockServiceClient, pairingClient pairingv1.PairingServiceClient) *server {
	return &server{lockClient: lockClient, pairingClient: pairingClient}
}

// Need to implement certificate based authentication
func Start() {

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("CLOUD")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.LoggingMiddlewareGrpc),
	)

	//Instantiate the clients
	lockClient := NewLockClient()
	pairingClient := NewPairingClient()

	dependencies := newServer(*lockClient, *pairingClient)
	//Inject dependencies into the server

	//Register the server
	lockv1.RegisterLockServiceServer(s, dependencies)
	pairingv1.RegisterPairingServiceServer(s, dependencies)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))
}
