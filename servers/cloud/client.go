package cloud

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func NewLockClient() *lockv1.LockServiceClient {
	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("BRIDGE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := lockv1.NewLockServiceClient(conn)
	return &client
}

func NewPairingClient() *pairingv1.PairingServiceClient {

	conn, err := grpc.DialContext(
		context.Background(),
		"dns:///0.0.0.0"+os.Getenv("BRIDGE"),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	client := pairingv1.NewPairingServiceClient(conn)
	return &client
}
