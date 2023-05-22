package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	lockv1 "github.com/nohns/proto/lock/v1"
)

func newLockClient(addr string) *lockv1.LockServiceClient {

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

	log.Println("Dialing lock service at", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+addr,
		/* grpc.WithBlock(), */
		grpc.WithTransportCredentials(tlsCredential),
	)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := lockv1.NewLockServiceClient(conn)
	return &client
}
