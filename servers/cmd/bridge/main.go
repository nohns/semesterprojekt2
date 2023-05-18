package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nohns/servers/bridge/domain"
	"github.com/nohns/servers/bridge/server"

	"github.com/nohns/servers/bridge/uart"
	"github.com/nohns/servers/pkg/certificatestore"
	"github.com/nohns/servers/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//no magic autoloading of .env file
	config.ReadEnvfile()

	// Read config into struct
	conf, err := config.LoadConfFromEnv()
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

	if err := certificatestore.Migrate(db); err != nil {
		log.Printf("error migrating database: %v", err)
		return
	}
	defer db.Close()

	log.Println("Starting certificateStore")
	//Create certificate store
	store := certificatestore.New(db)

	//uart layer
	log.Println("Starting UART")
	uart := uart.New()

	//Domain layer
	log.Println("Starting Domain")
	domain := domain.New(uart, store)

	//Check if the bridge is already paired with a cloud
	//If not establish a temp certificate and handle the root

	certificate, err := domain.PairCloud(conf.CloudGRPCURI)
	if err != nil {
		log.Fatalln("Error pairing with cloud exiting gracefully: ", err)
	}

	//Server layer
	server := server.New(conf, domain)

	log.Println("Starting Server")
	server.Start(certificate)

	//use the certificate provided to secure the GRPC server connection with https

	//server.StartGRPCServer(domain)

}
