package cloud

import (
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc/grpclog"

	"github.com/nohns/servers/pkg/middleware"

	lockv1 "github.com/nohns/proto/lock/v1"
	lockv1connect "github.com/nohns/proto/lock/v1/lockv1connect"
	pairingv1 "github.com/nohns/proto/pairing/v1"
	pairingv1connect "github.com/nohns/proto/pairing/v1/pairingv1connect"
)

/* pb "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1"
pbConnect "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1/phonev1connect" */

//Server responbilbe for communication with the react native phone app

// This struct should take in
type server struct {
	lockv1connect.UnimplementedLockServiceHandler
	pairingv1connect.UnimplementedPairingServiceHandler

	lockClient    lockv1.LockServiceClient
	pairingClient pairingv1.PairingServiceClient
}

func newServer(lockClient lockv1.LockServiceClient, pairingClient pairingv1.PairingServiceClient) *server {
	return &server{lockClient: lockClient, pairingClient: pairingClient}
}

// it might not be working because its unable to connect to anything at this point
func Start() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	logger := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(logger)
	mux := http.NewServeMux()

	//Instantiate the clients
	lockClient /* , lockConn */ := NewLockClient()
	pairingClient /* , pairingConn */ := NewPairingClient()

	//Open tunnel server
	//go NewTunnel(lockConn)
	//go NewTunnel(pairingConn)

	middleware.LoggingMiddleware(mux)

	cloud := newServer(*lockClient, *pairingClient)

	lockPath, lockHandler := lockv1connect.NewLockServiceHandler(cloud)
	pairingPath, pairingHandler := pairingv1connect.NewPairingServiceHandler(cloud)
	log.Println("Lock path: " + lockPath)

	//Register handlers with the mux
	mux.Handle(lockPath, lockHandler)
	mux.Handle(pairingPath, pairingHandler)

	log.Println("Cloud server listening on: " + os.Getenv("CLOUD"))
	http.ListenAndServe(
		os.Getenv("CLOUD"),
		// Use h2c so we can serve HTTP/2 without TLS.
		//TODO: We need to add the certificate handling here with TLS
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
