package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/nohns/servers/bridge/bluetooth"
	"github.com/nohns/servers/bridge/hw"
	"github.com/nohns/servers/bridge/sqlite"
)

type app struct {
	peripheral *bluetooth.Peripheral
}

// Boostrap the bridge app, initialzing all dependencies
func bootstrap() (*app, error) {
	conf, err := loadConfFromEnv()
	if err != nil {
		return nil, fmt.Errorf("could not load config: %v", err)
	}

	p, err := bluetooth.PreparePeripheral()
	if err != nil {
		return nil, fmt.Errorf("could not prepare ble peripheral: %v", err)
	}

	// Open DB and migrate schema
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", conf.DBPath))
	if err != nil {
		return nil, fmt.Errorf("could not open sqlite database: %v", err)
	}
	if err := sqlite.Migrate(db); err != nil {
		return nil, fmt.Errorf("could not migrate sqlite database: %v", err)
	}

	return &app{
		peripheral: p,
	}, nil
}

// Run the bridge app and listen for on different hardware interfaces (ex. start pair button)
func (a *app) run() error {
	// Listen for hardare changes and set handlers
	hw.HandlePairStartPress(a.onPairStartPress)
	go hw.Listen()

	return nil
}

func (a *app) onPairStartPress() {
	if err := a.peripheral.Advertise(); err != nil {
		log.Printf("failed to start advertising bluetooth peripheral: %v", err)
		return
	}

	log.Println("started advertising bluetooth peripheral successfully")
}
