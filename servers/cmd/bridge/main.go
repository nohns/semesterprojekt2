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

	// Read config into struct
	conf, err := config.LoadConfFromEnv()
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}

	//uart layer
	log.Println("Starting UART")
	uart := uart.New()

	//Domain layer
	log.Println("Starting Domain")
	domain := domain.New(uart)

	//Server layer
	server := server.New(*conf, domain)

	log.Println("Starting Server")
	server.Start()

}
