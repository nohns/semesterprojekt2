package server

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	lockv1 "github.com/nohns/proto/lock/v1"
	"github.com/nohns/servers/pkg/config"
	mw "github.com/nohns/servers/pkg/middleware"
)

// Should probaly revoke the certificate when the user logs out

type server struct {
	lockv1.UnimplementedLockServiceServer

	domain domain
	config *config.Config
}

func New(config *config.Config, domain domain) *server {
	return &server{config: config, domain: domain}
}

func (s *server) Start( /* certificate *tls.Certificate */ ) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	lis, err := net.Listen("tcp", s.config.BridgeGRPCURI)
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

		grpc.ChainUnaryInterceptor(mw.Timeout, mw.LoggingMiddlewareGrpc),
	)

	//Register the server
	lockv1.RegisterLockServiceServer(server, s)

	//Idk why the fuck this is needed
	reflection.Register(server)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", s.config.BridgeGRPCURI)
	log.Fatal(server.Serve(lis))

}
