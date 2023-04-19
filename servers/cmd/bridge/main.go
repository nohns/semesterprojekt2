package main

import (
	"log"

	"github.com/nohns/servers/bridge/domain"
	"github.com/nohns/servers/bridge/server"
	"github.com/nohns/servers/bridge/uart"
	"github.com/nohns/servers/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//no magic autoloading of .env file
	config.ReadEnvfile()

	// Read config
	/* conf, err := loadConfFromEnv()
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	} */

	// Open database
	/* db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", conf.DBPath))
	if err != nil {
		log.Printf("error opening database: %v", err)
		return
	}
	if err := sqlite.Migrate(db); err != nil {
		log.Printf("error migrating database: %v", err)
		return
	}
	defer db.Close() */

	// Create event store
	//var store eventsource.EventStore = sqlite.NewEventSource(db)

	//Repository layer
	//repo := repository.New()

	//uart layer
	log.Println("Starting UART")
	uart := uart.New()

	//Domain layer
	log.Println("Starting Domain")
	domain := domain.New(uart)

	//Server layer
	log.Println("Starting Server")
	server.StartGRPCServer(domain)

}
