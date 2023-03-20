package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
	"github.com/nohns/servers/bridge"
	"github.com/nohns/servers/bridge/cmdstream"
	"github.com/nohns/servers/pkg/eventbus"
	"github.com/nohns/servers/pkg/eventsource"
	"github.com/nohns/servers/pkg/sqlite"
	"google.golang.org/grpc"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//no magic autoloading of .env file
	readEnvfile()

	// Read config
	conf, err := loadConfFromEnv()
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}

	// Open database
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", conf.DBPath))
	if err != nil {
		log.Printf("error opening database: %v", err)
		return
	}
	if err := sqlite.Migrate(db); err != nil {
		log.Printf("error migrating database: %v", err)
		return
	}
	defer db.Close()

	// Create event store
	var store eventsource.EventStore = sqlite.NewEventSource(db)

	// Create event bus
	evtbus := eventbus.New()

	// Create lock domain service
	lockService := bridge.NewLockService(store, evtbus)

	conn, err := grpc.Dial(conf.CloudGRPCURI, grpc.WithInsecure())
	if err != nil {
		log.Printf("error dialing cloud grpc: %v", err)
		return
	}
	defer conn.Close()

	// Create command stream
	distributor := cmdstream.NewDistributor(lockService)
	cs := cmdstream.New(bridgepb.NewCmdServiceClient(conn), distributor)

	// Start listening for commands
	if err := cs.Listen(context.TODO()); err != nil {
		log.Printf("error streaming commands: %v", err)
		return
	}
}
