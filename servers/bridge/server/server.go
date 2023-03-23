package server

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	pairingv1 "github.com/nohns/proto/pairing/v1"
	mw "github.com/nohns/servers/pkg/middleware"
)

type Server struct {
pairingv1.UnimplementedPairingServiceServer
}

func newServer() *Server {
	return &Server{}
}

func StartRESTServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/v1/certificate", certificateHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln(err)
	}

}

func StartGRPCServer() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	addr := os.Getenv("GRPC")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(mw.LoggingMiddlewareGrpc),
	)
	//Inject dependencies into the server
	dependencies := newServer()

	//Register the server
	pairingv1.RegisterPairingServiceServer(s, dependencies)

	//Idk why the fuck this is needed
	reflection.Register(s)

	// Serve gRPC Server
	log.Info("Serving gRPC on http://", addr)
	log.Fatal(s.Serve(lis))

}





// Should probaly revoke the certificate when the user logs out