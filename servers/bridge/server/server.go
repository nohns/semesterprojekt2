package server

import (
	"crypto/tls"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
	"github.com/nohns/servers/pkg/config"
	mw "github.com/nohns/servers/pkg/middleware"
)

// Should probaly revoke the certificate when the user logs out

type server struct {
	lockv1.UnimplementedLockServiceServer
	pairingv1.UnimplementedPairingServiceServer

	domain domain
	config *config.Config
}

func New(config *config.Config, domain domain) *server {
	return &server{config: config, domain: domain}
}

func (s *server) Start(certificate *tls.Certificate) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", s.config.BridgeGRPCURI)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		Certificates:       []tls.Certificate{*certificate},
		ClientCAs:          nil,
		ClientAuth:         tls.RequireAndVerifyClientCert,
	}

	server := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.ChainUnaryInterceptor(mw.Timeout, mw.LoggingMiddlewareGrpc),
	)

	//Register the server
	pairingv1.RegisterPairingServiceServer(server, s)
	lockv1.RegisterLockServiceServer(server, s)

	//Idk why the fuck this is needed
	reflection.Register(server)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", s.config.BridgeGRPCURI)
	log.Fatal(server.Serve(lis))

}
