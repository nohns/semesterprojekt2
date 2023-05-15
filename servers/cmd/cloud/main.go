package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nohns/servers/cloud/domain"
	"github.com/nohns/servers/cloud/server"
	"github.com/nohns/servers/pkg/certificatestore"
	"github.com/nohns/servers/pkg/config"
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

	server := server.New(conf, domain)

	//Instantiate the connect server and the bridge clients
	server.Start([]byte{})
}

//the config contains
//DBPath string
//server port of the cloud server
//bridge port of the bridge server
