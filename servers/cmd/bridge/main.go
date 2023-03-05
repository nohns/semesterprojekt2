package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/nohns/semesterprojekt2/bridge"
	"github.com/nohns/semesterprojekt2/bridge/cmdstream"
	"github.com/nohns/semesterprojekt2/pkg/eventsource"
	"github.com/nohns/semesterprojekt2/pkg/sqlite"
	bridgepb "github.com/nohns/semesterprojekt2/proto/gen/go/cloud/bridge/v1"
	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

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

	// Create lock service
	lockService := bridge.NewLockService(store, nil)

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
		log.Printf("command stream closed")
		return
	}
}
