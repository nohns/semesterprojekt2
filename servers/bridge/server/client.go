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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	//instead we pass in the certificate
	/* tlsCert, err := makeTLSCert(username)
	if err != nil {
		log.Fatalln("Failed to create client cert:", err)
	} */

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*tlsCert},
		RootCAs:      nil,
	}

	conn, err := grpc.DialContext(ctx, "dns:///0.0.0.0"+addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()
	c := pairingv1.NewPairingServiceClient(conn)
	return c, conn
}
