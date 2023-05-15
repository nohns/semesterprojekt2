package server

import (
	"crypto/tls"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/nohns/servers/pkg/config"
	"github.com/nohns/servers/pkg/middleware"

	lockv1 "github.com/nohns/proto/lock/v1"

	pairingv1 "github.com/nohns/proto/pairing/v1"
)

type Pairing interface {
}

// This struct should take in
type server struct {
	lockv1.UnimplementedLockServiceServer
	pairingv1.UnimplementedPairingServiceServer

	lockClient    lockv1.LockServiceClient
	pairingClient pairingv1.PairingServiceClient
	pairing       Pairing

	config *config.Config
}

func New(c *config.Config, p Pairing) *server {

	//Open the client connections
	lockClient := newLockClient(c.BridgeGRPCURI)
	pairingClient := newPairingClient(c.BridgeGRPCURI)

	return &server{config: c, pairing: p, lockClient: *lockClient, pairingClient: *pairingClient}
}

// Need to implement certificate based authentication
func (s *server) Start(csr []byte) {

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", s.config.CloudGRPCURI)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{},
		ClientCAs:          nil,
		ClientAuth:         tls.RequireAndVerifyClientCert,
	}

	server := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.UnaryInterceptor(middleware.LoggingMiddlewareGrpc),
	)

	//Register the server
	lockv1.RegisterLockServiceServer(server, s)
	pairingv1.RegisterPairingServiceServer(server, s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", s.config.CloudGRPCURI)
	log.Fatal(server.Serve(lis))
}
