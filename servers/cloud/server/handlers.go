package server

import (
	"context"
	"errors"
	"log"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

func (s *server) GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error) {

	// check for required field
	if req.Id == "" {
		return nil, errors.New("id is required")
	}
	log.Println("req: ", req)

	// call lock service
	res, err := s.lockClient.GetLockState(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	//Return the response
	return res, nil

}

func (s *server) SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error) {

	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	// call lock service
	res, err := s.lockClient.SetLockState(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	//Return the response

	return res, nil
}

func (s *server) Register(ctx context.Context, req *pairingv1.RegisterRequest) (*pairingv1.RegisterResponse, error) {
	log.Println("req: ", req)
	if req.Csr == nil || req.PublicKey == nil {
		return nil, errors.New("id is required")
	}

	// call pairing service
	res, err := s.pairingClient.Register(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	//Return the response

	return res, nil
}
