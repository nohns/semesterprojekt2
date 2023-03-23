package cloud

import (
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	lockv1 "github.com/nohns/proto/lock/v1"
	lockv1connect "github.com/nohns/proto/lock/v1/lockv1connect"
	pairingv1 "github.com/nohns/proto/pairing/v1"
	pairingv1connect "github.com/nohns/proto/pairing/v1/pairingv1connect"
)

/* pb "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1"
pbConnect "github.com/nohns/semesterprojekt2/servers/cloud/phone/v1/phonev1connect" */

//Server responbilbe for communication with the react native phone app

//This struct should take in
type server struct {
lockv1connect.UnimplementedLockServiceHandler
pairingv1connect.UnimplementedPairingServiceHandler
lockClient lockv1.LockServiceClient
pairingClient pairingv1.PairingServiceClient
}

func newServer(lockClient lockv1.LockServiceClient, pairingClient pairingv1.PairingServiceClient) *server {
    return &server{lockClient: lockClient, pairingClient: pairingClient}
}

func Start() {
    mux := http.NewServeMux()

    //Instantiate the clients
    lockClient := NewLockClient()
    pairingClient := NewPairingClient()
    

	cloud := newServer(*lockClient, *pairingClient)

    lockPath, lockHandler := lockv1connect.NewLockServiceHandler(cloud)
    pairingPath, pairingHandler := pairingv1connect.NewPairingServiceHandler(cloud)

//Register handlers with the mux
    mux.Handle(lockPath, lockHandler)
    mux.Handle(pairingPath, pairingHandler)

    http.ListenAndServe(
        os.Getenv("CLOUD"),
        // Use h2c so we can serve HTTP/2 without TLS.
        //TODO: We need to add the certificate handling here with TLS
        h2c.NewHandler(mux, &http2.Server{}),
    )
}

