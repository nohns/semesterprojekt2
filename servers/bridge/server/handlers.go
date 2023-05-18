package server

import (
	"context"
	"log"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
	"google.golang.org/grpc/status"
)

// Take whatever functions the domain needs to use should be injected into this interface
type domain interface {
	Register() (string, error)
	GetLock(ctx context.Context) (bool, error)
	SetLock(ctx context.Context, state bool) (bool, error)
}

func (s *server) Register(ctx context.Context, req *pairingv1.RegisterRequest) (*pairingv1.RegisterResponse, error) {
	//Veryfiy that the certificate is valid
	if req.Csr == nil {
		log.Println("CSR is nil")
		return nil, status.Error(400, "CSR is nil")
	}

	//Call buissness logic

	//Return response
	return &pairingv1.RegisterResponse{
		BridgeId: "TEMPVALUE",
		Cert:     nil,
	}, nil
}

func (s *server) GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error) {
	log.Println("GetLock called")
	//veryify that id is not empty
	if req.Id == "" {
		log.Println("ID is empty")
		return nil, status.Error(400, "ID is empty")
	}

	//Call buissness logic
	state, err := s.domain.GetLock(ctx)
	if err != nil {
		log.Println("Error getting lock state: ", err)
		return nil, status.Error(500, "Error getting lock state")
	}

	//Return response
	return &lockv1.GetLockStateResponse{
		Engaged: state,
	}, nil
}

func (s *server) SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error) {
	log.Println("SetLock called")
	//veryify that id is not empty
	if req.Id == "" {
		log.Println("ID is empty")
		return nil, status.Error(400, "ID is empty")
	}

	//Call buissness logic
	state, err := s.domain.SetLock(ctx, req.Engaged)
	if err != nil {
		log.Println("Error setting lock state: ", err)
		return nil, status.Error(500, "Error setting lock state")
	}

	//Return response
	return &lockv1.SetLockStateResponse{
		Engaged: state,
	}, nil
}
