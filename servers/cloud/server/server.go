package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nohns/servers/pkg/config"
	"github.com/nohns/servers/pkg/middleware"

	lockv1 "github.com/nohns/proto/lock/v1"
)

// This struct should take in
type server struct {
	lockv1.UnimplementedLockServiceServer
	controller controller

	config config.Config
}

func New(c config.Config) *server {

	//Open the client connections

	return &server{config: c}
}

// Need to implement certificate based authentication
func (s *server) Start() {

	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", s.config.CloudGRPCURI)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	caPem, err := ioutil.ReadFile("../../../cert/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		log.Fatal(err)
	}

	serverCert, err := tls.LoadX509KeyPair("../../../cert/server-cert.pem", "../../../cert/server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(conf)

	server := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(middleware.LoggingMiddlewareGrpc),
	)

	//Register the server
	lockv1.RegisterLockServiceServer(server, s)

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

	caCert, err := ioutil.ReadFile("../../../cert/ca-cert.pem")
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("../../../cert/client-cert.pem", "../../../cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// set config of tls credential
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	//The reverse proxy connects to the GRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		/* "dns:///0.0.0.0:8080", */
		"dns:///0.0.0.0"+os.Getenv("CLOUD_GRPC_URI"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(tlsCredential),
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
