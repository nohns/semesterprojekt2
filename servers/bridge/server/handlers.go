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
	GetLock() (bool, error)
	SetLock() (bool, error)
}

func (s *Server) Register(ctx context.Context, req *pairingv1.RegisterRequest) (*pairingv1.RegisterResponse, error) {
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

func (s *Server) GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error) {
	log.Println("GetLock called")
	//veryify that id is not empty
	if req.Id == "" {
		log.Println("ID is empty")
		return nil, status.Error(400, "ID is empty")
	}

	//Call buissness logic
	s.domain.GetLock()

	//Return response
	return &lockv1.GetLockStateResponse{
		Locked: true,
	}, nil
}

func (s *Server) SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error) {
	log.Println("SetLock called")
	//veryify that id is not empty
	if req.Id == "" {
		log.Println("ID is empty")
		return nil, status.Error(400, "ID is empty")
	}

	//Call buissness logic

	//Return response
	return &lockv1.SetLockStateResponse{
		Locked: true,
	}, nil
}
