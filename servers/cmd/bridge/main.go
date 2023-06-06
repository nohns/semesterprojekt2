package main

import (
	"log"

	"github.com/nohns/servers/bridge/bluetooth"
	"github.com/nohns/servers/bridge/domain"
	"github.com/nohns/servers/bridge/hw"
	"github.com/nohns/servers/bridge/uart"
	"github.com/nohns/servers/cloud/server"

	"github.com/nohns/servers/pkg/config"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	//no magic autoloading of .env file
	config.ReadEnvfile()

	// Read config into struct
	conf, err := config.LoadConfFromEnv()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// UART layer
	log.Println("Starting UART")
	uart := uart.New()

	// Domain layer
	log.Println("Starting Domain")
	domain := domain.New(uart)

	// Start hardware layer with bluetooth and button
	blePeriph, err := bluetooth.PreparePeripheral(domain)
	if err != nil {
		log.Fatalf("could not prepare ble peripheral: %v", err)
	}
	bh := hw.NewButtonHandler(blePeriph)
	hw := hw.New(bh)

	//Server layer
	_ = server.New(*conf)

	log.Println("Starting HW button listener")
	hw.Listen()

	log.Println("Starting Server")
	//server.Start()
}
