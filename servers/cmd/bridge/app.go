package main

import (
	"fmt"
	"log"

	"github.com/nohns/servers/bridge/bluetooth"
	"github.com/nohns/servers/bridge/hw"
)

type app struct {
	peripheral *bluetooth.Peripheral
}

// Boostrap the bridge app, initialzing all dependencies
func bootstrap() (*app, error) {

	p, err := bluetooth.PreparePeripheral()
	if err != nil {
		return nil, fmt.Errorf("could not prepare ble peripheral: %v", err)
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
