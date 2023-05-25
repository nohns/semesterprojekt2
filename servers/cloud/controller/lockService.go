package controller

import (
	"context"

	lockv1 "github.com/nohns/proto/lock/v1"
	"github.com/nohns/servers/cloud/client"
	"github.com/nohns/servers/pkg/config"
)

type controller struct {
	lockClient lockv1.LockServiceClient
}

func New(c config.Config) *controller {
	lockClient := client.NewLockClient(c.BridgeGRPCURI)

	return &controller{lockClient: *lockClient}
}

func (c *controller) GetLockState(ctx context.Context, req *lockv1.GetLockStateRequest) (*lockv1.GetLockStateResponse, error) {
	return c.lockClient.GetLockState(ctx, req)
}

func (c *controller) SetLockState(ctx context.Context, req *lockv1.SetLockStateRequest) (*lockv1.SetLockStateResponse, error) {
	return c.lockClient.SetLockState(ctx, req)
}
