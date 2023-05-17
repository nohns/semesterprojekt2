package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nohns/servers/cloud/domain"
	"github.com/nohns/servers/cloud/server"
	"github.com/nohns/servers/pkg/certificatestore"
	"github.com/nohns/servers/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	config.ReadEnvfile()

	//Open database connection
	conf, err := config.LoadConfFromEnv()
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}

	// Open database
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

	//Create certificate store which implements the repository interface
	cerStore := certificatestore.New(db)

	domain := domain.New(cerStore)

	//Create self signed root certificate if it doesn't exit
	rootCertificate, err := domain.InitializeRootCertificate()
	if err != nil {
		log.Fatalln("Error initializing root certificate: ", err)
	}

	log.Println("Constructing server object")
	server := server.New(conf, domain)

	log.Println("Starting the server")
	//Instantiate the connect server and the bridge clients
	server.Start(rootCertificate)

	//Here we prolly want to await a signing request from the bridge
}

//the config contains
//DBPath string
//server port of the cloud server
//bridge port of the bridge server
