package server

import (
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
	mw "github.com/nohns/servers/pkg/middleware"
)

type Server struct {
	pairingv1.UnimplementedPairingServiceServer
	lockv1.UnimplementedLockServiceServer
	domain domain
}

func newServer(domain domain) *Server {
	return &Server{domain: domain}
}

// Inject domain layer into the server
func StartGRPCServer(domain domain) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("BRIDGE")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(mw.LoggingMiddlewareGrpc),
	)

	dependencies := newServer(domain)

	//Create tunnel
	ServeTunnel(s, dependencies)

	//Inject dependencies into the server

	//Register the server
	pairingv1.RegisterPairingServiceServer(s, dependencies)
	lockv1.RegisterLockServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}

// Should probaly revoke the certificate when the user logs out
