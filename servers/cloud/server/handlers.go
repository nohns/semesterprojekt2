package server

import (
	"context"
	"errors"
	"log"

	lockv1 "github.com/nohns/proto/lock/v1"
)

type controller interface {
	GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error)
	SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error)
}

func (s *server) GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error) {
	log.Println("req: ", req)
	// check for required field
	if req.Id == "" {
		return nil, errors.New("id is required")
	}
	log.Println("req: ", req)

	// call lock service
	res, err := s.controller.GetLockState(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	//Return the response
	return res, nil

}

func (s *server) SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error) {
	log.Println("req: ", req)

	if req.Id == "" {
		return nil, errors.New("id is required")
	}

	// call lock service
	res, err := s.controller.SetLockState(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	//Return the response

	return res, nil
}
