package cmdstream

import (
	"context"

	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
	"github.com/nohns/servers/bridge"
)

type distributor struct {
	lockSvc *bridge.LockService
}

func NewDistributor(lockSvc *bridge.LockService) *distributor {
	return &distributor{
		lockSvc: lockSvc,
	}
}

// Distributes commands to the correct services
func (d *distributor) distribute(ctx context.Context, streamcmd *bridgepb.StreamCommand) error {
	if err := d.distributeLockCmds(ctx, streamcmd); err != nil {
		return err
	}

	return nil
}
