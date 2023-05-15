package bluetooth

import (
	"bytes"
	"fmt"
	"log"

	"github.com/nohns/servers/bridge/certificate"
	"tinygo.org/x/bluetooth"
)

var (
	advertisementLocalName = "Smart Lock"
	serviceUUID            = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x10} // "9b7155fc-d47e-4309-9c81-a2261d582810"
	characteristicCsrUUID  = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x11} // "9b7155fc-d47e-4309-9c81-a2261d582811"
	characteristicCertUUID = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x12} //"9b7155fc-d47e-4309-9c81-a2261d582812"
)

type Peripheral struct {
	adv             *bluetooth.Advertisement
	svc             *bluetooth.Service
	csrChar         *bluetooth.Characteristic
	certChar        *bluetooth.Characteristic
	connRecvBuffers map[bluetooth.Connection]*bytes.Buffer
}

// PreparePeripheral prepares the bluetooth peripheral by enabling the bluetooth adapter underneeth
// and configuring the advertisement information.
func PreparePeripheral() (*Peripheral, error) {

	// Enable the use of the OS supported bluetooth adapter. BlueZ for Linux running on a Raspberry Pi.
	adapter := bluetooth.DefaultAdapter
	if err := adapter.Enable(); err != nil {
		return nil, fmt.Errorf("failed to enable ble adapter: %v", err)
	}

	// Configure advertisement with our own uuid and name.
	adv := adapter.DefaultAdvertisement()
	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    advertisementLocalName,
		ServiceUUIDs: []bluetooth.UUID{bluetooth.NewUUID(serviceUUID)},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to configure advertisement: %v", err)
	}

	// Set up the bluetooth peripheral.
	p := &Peripheral{
		adv:             adv,
		csrChar:         new(bluetooth.Characteristic),
		certChar:        new(bluetooth.Characteristic),
		connRecvBuffers: make(map[bluetooth.Connection]*bytes.Buffer),
	}

	// Define the bluetooth service and characteristics of the Smart lock peripheral.
	p.svc = &bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle:     p.csrChar,
				UUID:       bluetooth.NewUUID(characteristicCsrUUID),
				Flags:      bluetooth.CharacteristicWritePermission | bluetooth.CharacteristicWriteWithoutResponsePermission,
				WriteEvent: p.handleCSRWrite,
			},
			{
				Handle: p.certChar,
				UUID:   bluetooth.NewUUID(characteristicCertUUID),
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
		},
	}

	return p, nil
}

// Advertise starts the advertisement of the bluetooth peripheral, which makes discoverable to the smartphone app
func (p *Peripheral) Advertise() error {
	if err := p.adv.Start(); err != nil {
		return fmt.Errorf("could not start ble advertisement: %v", err)
	}

	return nil
}

// StopAdvertise stops the advertisement of the bluetooth peripheral, which makes it undiscoverable to the smartphone app
func (p *Peripheral) StopAdvertise() error {
	if err := p.adv.Stop(); err != nil {
		return fmt.Errorf("could not stop ble advertisement: %v", err)
	}

	return nil
}

// handleCSR handles the fully received CSR data from the smartphone app.
func (p *Peripheral) handleCSR(pemdata []byte) error {
	phoneCSR, err := certificate.ParseCSR(pemdata)
	if err != nil {
		return fmt.Errorf("could not parse csr data received: %v", err)
	}

	return nil
}

// handleCSRWrite handles the write event of the CSR ble characteristic.
func (p *Peripheral) handleCSRWrite(client bluetooth.Connection, int, value []byte) {
	// Get the current CSR buffer for the client connection.
	b, ok := p.connRecvBuffers[client]
	if !ok {
		p.connRecvBuffers[client] = &bytes.Buffer{}
	}

	// If a zero byte is encountered, then handle the buffered CSR data and reset the recv buffer
	if len(value) == 1 && value[0] == 0 {
		p.handleCSR(b.Bytes())
		b.Reset()
		return
	}

	// Write the received data to the buffer
	if _, err := b.Write(value); err != nil {
		log.Printf("failed to write csr data to buffer: %v", err)
		return
	}
}

// writeCert takes an arbitrary number of bytes in a slice and send them as chucks of 128 bytes to
// the smartphone app.
func (p *Peripheral) writeCert(data []byte) error {
	// Send the data in chunks of 128 bytes.
	for i := 0; i < len(data); i += 128 {
		end := i + 128
		if end > len(data) {
			end = len(data)
		}

		if _, err := p.certChar.Write(data[i:end]); err != nil {
			return fmt.Errorf("failed to write cert data: %v", err)
		}
	}

	// Send a zero byte to indicate the end of the data.
	if _, err := p.certChar.Write([]byte{0}); err != nil {
		return fmt.Errorf("failed to write cert data: %v", err)
	}
	return nil
}
