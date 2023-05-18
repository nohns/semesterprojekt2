package server

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func RunWithTLSAuth(addr string, tlsCert *tls.Certificate) (pairingv1.PairingServiceClient, *grpc.ClientConn) {
	log.Println("Dialing pairing service at", addr)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*tlsCert},
	}

	conn, err := grpc.DialContext(ctx, "dns:///0.0.0.0"+addr,

		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()
	c := pairingv1.NewPairingServiceClient(conn)
	log.Println("Connected to pairing service")
	return c, conn
}
