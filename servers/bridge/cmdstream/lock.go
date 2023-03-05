package cmdstream

import (
	"context"

	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
	"google.golang.org/protobuf/proto"
)

func (d *distributor) distributeLockCmds(ctx context.Context, streamcmd *bridgepb.StreamCommand) error {
	switch streamcmd.Type {
	case "DisengageLock":
		var cmd bridgepb.StreamCommand_DisengageLock
		if err := proto.Unmarshal(streamcmd.Cmd, &cmd); err != nil {
			return err
		}

		return d.lockSvc.Disengage(ctx, cmd.LockId)
	case "EngageLock":
		var cmd bridgepb.StreamCommand_EngageLock
		if err := proto.Unmarshal(streamcmd.Cmd, &cmd); err != nil {
			return err
		}

		return d.lockSvc.Engage(ctx, cmd.LockId)
	}

	return nil
}
