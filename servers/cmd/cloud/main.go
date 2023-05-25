package main

import (
	"log"

	"github.com/nohns/servers/cloud/server"
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

	log.Println("Constructing server object")
	server := server.New(*conf)

	log.Println("Starting the server")
	//Instantiate the connect server and the bridge clients
	server.Start()

	//Here we prolly want to await a signing request from the bridge
}
