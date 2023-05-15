package server

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func newLockClient(addr string) *lockv1.LockServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := lockv1.NewLockServiceClient(conn)
	return &client
}

func newPairingClient(addr string) *pairingv1.PairingServiceClient {

	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+addr,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := pairingv1.NewPairingServiceClient(conn)
	return &client
}
