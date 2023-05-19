package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nohns/servers/pkg/config"
	"github.com/nohns/servers/pkg/middleware"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

type Pairing interface {
	PairCloud(csrPEM []byte) ([]byte, error)
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
func (s *server) Start( /* rootCertificate *tls.Certificate */ ) {

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", s.config.CloudGRPCURI)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	/* tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{*rootCertificate},
		RootCAs:            nil, //Set to NIL because cloud itself is the root CA
		ClientAuth:         tls.RequireAnyClientCert,
		ClientCAs:          certPool,
	} */

	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(middleware.LoggingMiddlewareGrpc),
	)

	//Register the server
	lockv1.RegisterLockServiceServer(server, s)
	pairingv1.RegisterPairingServiceServer(server, s)

	// Serve gRPC Server

	go func() {
		log.Info("Serving gRPC on http://", s.config.CloudGRPCURI)
		log.Fatal(server.Serve(lis))
	}()

	err = startGateway()
	log.Fatalln("Failed to start gateway:", err)
}

//Ok the major issue is the fact that the cloud server has TLS settings on that require a certificate
//The bridge server does not have a valid certirifacrte
//I think the solution is to open up a temporary server with lesser TLS settings which should be closed after the pairing is complete

// run the generated GRPC gateway server
func startGateway() error {
	log := grpclog.NewLoggerV2WithVerbosity(os.Stdout, io.Discard, io.Discard, 1)
	grpclog.SetLoggerV2(log)

	//The reverse proxy connects to the GRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+os.Getenv("CLOUD_GRPC_URI"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	//Main mux where options are added in
	gwmux := runtime.NewServeMux()

	if err = lockv1.RegisterLockServiceHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %v", err)
	}

	gatewayAddress := os.Getenv("CLOUD_GRPC_GATEWAY_URI")

	//middleware chaining
	middleware := middleware.LoggingMiddleware(middleware.Cors(gwmux))

	gwServer := &http.Server{
		Addr:    gatewayAddress,
		Handler: middleware,
	}

	log.Info("Serving gRPC-Gateway", gatewayAddress)
	log.Fatalln(gwServer.ListenAndServe())

	return nil
}
