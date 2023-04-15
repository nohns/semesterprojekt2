package cloud

import (
	"context"
	"errors"
	"log"

	"github.com/bufbuild/connect-go"

	lockv1 "github.com/nohns/proto/lock/v1"
	pairingv1 "github.com/nohns/proto/pairing/v1"
)

// This function is used to get the lock state of a lock.
// It calls the GetLockState function of the lock service.
// It uses the lock id passed in the request to get the lock state.
// It returns a response containing the lock state.

func (s *server) GetLockState(ctx context.Context, in *connect.Request[lockv1.GetLockStateRequest]) (*connect.Response[lockv1.GetLockStateResponse], error) {
	log.Println("GetLockState called")
	// check for required field
	if in.Msg.Id == "" {
		return nil, errors.New("id is required")
	}

	// create request for lock service
	req := &lockv1.GetLockStateRequest{
		Id: in.Msg.Id,
	}
	log.Println("req: ", req)

	// call lock service
	res, err := s.lockClient.GetLockState(ctx, req)
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	log.Println("res: ", res)

	// convert response to connect response
	resp := &connect.Response[lockv1.GetLockStateResponse]{
		Msg: &lockv1.GetLockStateResponse{
			Locked: true,
		},
	}

	return resp, nil
}

// SetLockState is a handler for the LockService.SetLockState RPC method.
// It accepts a lockv1.SetLockStateRequest as a connect.Request and
// returns a lockv1.SetLockStateResponse as a connect.Response.
func (s *server) SetLockState(ctx context.Context, in *connect.Request[lockv1.SetLockStateRequest]) (*connect.Response[lockv1.SetLockStateResponse], error) {
	log.Println("SetLockState called")
	// Create a request to the lock service.
	/* req := &lockv1.SetLockStateRequest{
		Id:     in.Msg.Id,
		Locked: in.Msg.Locked,
	} */

	// Call the lock service.
	/* res, err := s.lockClient.SetLockState(ctx, req)
	if err != nil {
		return nil, err
	}

	// Create a response to the caller.
	resp := &connect.Response[lockv1.SetLockStateResponse]{
		Msg: res,
	}

	_ = resp

	//create fake response
	resp = &connect.Response[lockv1.SetLockStateResponse]{
		Msg: &lockv1.SetLockStateResponse{
			Locked: true,
		},
	} */

	return nil, nil
}

// Register is a unary RPC that receives a CSR and returns a certificate
// pair for the device.
// The CSR is validated to ensure that it is a valid CSR.
// The CSR is signed using the device CA, which is generated on the first
// request.
func (s *server) Register(ctx context.Context, in *connect.Request[pairingv1.RegisterRequest]) (*connect.Response[pairingv1.RegisterResponse], error) {
	log.Println("Register called")

	//validate its not empty
	/* 	if in.Msg.Csr == nil {
	   		return nil, errors.New("name is required")
	   	}


	   	//create request for pairing service
	   	req := &pairingv1.RegisterRequest{
	   		Csr: in.Msg.Csr,
	   	}

	   	//call pairing service
	   	res, err := s.pairingClient.Register(ctx, req)
	   	if err != nil {
	   		return nil, err
	   	}

	   	//convert response to connect response
	   	resp := &connect.Response[pairingv1.RegisterResponse]{
	   		Msg: res,
	   	} */

	return nil, nil
}
