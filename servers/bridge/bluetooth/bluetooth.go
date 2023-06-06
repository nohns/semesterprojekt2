package bluetooth

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"sync"
	"time"

	"tinygo.org/x/bluetooth"
)

var (
	advertisementLocalName    = "Smart Lock Bridge (beta)"
	serviceUUID               = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x10} // "9b7155fc-d47e-4309-9c81-a2261d582810"
	characteristicCSRUUID     = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x11} // "9b7155fc-d47e-4309-9c81-a2261d582811"
	characteristicCertUUID    = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x12} //"9b7155fc-d47e-4309-9c81-a2261d582812"
	characteristicCSRLenUUID  = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x13} // "9b7155fc-d47e-4309-9c81-a2261d582813"
	characteristicCertLenUUID = [16]byte{0x9b, 0x71, 0x55, 0xfc, 0xd4, 0x7e, 0x43, 0x09, 0x9c, 0x81, 0xa2, 0x26, 0x1d, 0x58, 0x28, 0x14} // "9b7155fc-d47e-4309-9c81-a2261d582814"
)

type recvBuf struct {
	bytes.Buffer
	len     int
	recvLen int
}

type signer interface {
	SignCertificate(csrReq []byte) ([]byte, error)
}

type Peripheral struct {
	adv                *bluetooth.Advertisement
	svc                *bluetooth.Service
	csrChar            *bluetooth.Characteristic
	certChar           *bluetooth.Characteristic
	lenCSRChar         *bluetooth.Characteristic
	lenCertChar        *bluetooth.Characteristic
	connRecvBuffers    map[bluetooth.Connection]*recvBuf
	mu                 sync.Mutex
	handshakeListeners []chan struct{}
	signer             signer
}

// PreparePeripheral prepares the bluetooth peripheral by enabling the bluetooth adapter underneeth
// and configuring the advertisement information.
func PreparePeripheral(signer signer) (*Peripheral, error) {

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
		adv:                adv,
		csrChar:            new(bluetooth.Characteristic),
		certChar:           new(bluetooth.Characteristic),
		lenCSRChar:         new(bluetooth.Characteristic),
		lenCertChar:        new(bluetooth.Characteristic),
		connRecvBuffers:    make(map[bluetooth.Connection]*recvBuf),
		handshakeListeners: make([]chan struct{}, 0),
		signer:             signer,
	}

	// Define the bluetooth service and characteristics of the Smart lock peripheral.
	p.svc = &bluetooth.Service{
		UUID: bluetooth.NewUUID(serviceUUID),
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle:     p.csrChar,
				UUID:       bluetooth.NewUUID(characteristicCSRUUID),
				Flags:      bluetooth.CharacteristicWritePermission | bluetooth.CharacteristicWriteWithoutResponsePermission,
				WriteEvent: p.handleCSRWrite,
			},
			{
				Handle: p.certChar,
				UUID:   bluetooth.NewUUID(characteristicCertUUID),
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
			{
				Handle:     p.lenCSRChar,
				UUID:       bluetooth.NewUUID(characteristicCSRLenUUID),
				Flags:      bluetooth.CharacteristicWritePermission | bluetooth.CharacteristicWriteWithoutResponsePermission,
				WriteEvent: p.handleCSRLenWrite,
			},
			{
				Handle: p.lenCertChar,
				UUID:   bluetooth.NewUUID(characteristicCertLenUUID),
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
		},
	}
	if err := adapter.AddService(p.svc); err != nil {
		return nil, fmt.Errorf("failed to add ble service: %v", err)
	}

	return p, nil
}

// BeginHandshake starts the advertisement of the bluetooth peripheral, which makes discoverable to the smartphone app
func (p *Peripheral) BeginHandshake(ctx context.Context) error {
	p.mu.Lock()

	log.Printf("starting ble advertisement")
	if err := p.adv.Start(); err != nil {
		return fmt.Errorf("could not start ble advertisement: %v", err)
	}
	p.mu.Unlock()

	go func() {
		log.Printf("waiting for ble handshake context to complete")
		<-ctx.Done()
		p.stopAdvertise()
		p.closeHandshakeListeners()
	}()

	// wait for the handshake to complete
	p.awaitHandshake()

	return nil
}

// awaitHandshake blocks until the handshake with the smartphone app has been completed.
func (p *Peripheral) awaitHandshake() {
	p.mu.Lock()
	c := make(chan struct{})
	p.handshakeListeners = append(p.handshakeListeners, c)
	p.mu.Unlock()

	log.Printf("awaiting ble handshake to complete")
	<-c
	log.Printf("ble handshake completed")
}

func (p *Peripheral) closeHandshakeListeners() {
	for _, b := range p.handshakeListeners {
		close(b)
	}
	p.handshakeListeners = make([]chan struct{}, 0)
}

// stopAdvertise stops the advertisement of the bluetooth peripheral, which makes it undiscoverable to the smartphone app
func (p *Peripheral) stopAdvertise() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.adv.Stop(); err != nil {
		return fmt.Errorf("could not stop ble advertisement: %v", err)
	}

	return nil
}

// handleCSR handles the fully received CSR data from the smartphone app.
func (p *Peripheral) performHandshake(pemdata []byte) error {
	// Sign the CSR data received from the smartphone app.
	cert, err := p.signer.SignCertificate(pemdata)
	if err != nil {
		return fmt.Errorf("could not sign csr data received: %v", err)
	}

	// Send the signed certificate to the smartphone app.
	p.writeCert(cert)

	// Signal successful handshake and reset the listeners
	p.closeHandshakeListeners()

	return nil
}

// handleCSRLenWrite handles the write event of the CSR length ble characteristic. This means that
// we receive the length of the CSR data that will be sent to us.
func (p *Peripheral) handleCSRLenWrite(client bluetooth.Connection, _ int, data []byte) {
	// Initialize receive buffer for the client connection, with the length of the CSR data given as data
	log.Printf("raw csr len data: %x", data[:2])
	len := int(binary.BigEndian.Uint16(data[:2]))
	log.Printf("recv csr len: %v", len)
	p.connRecvBuffers[client] = &recvBuf{
		len: len,
	}
}

// handleCSRWrite handles the write event of the CSR ble characteristic.
func (p *Peripheral) handleCSRWrite(client bluetooth.Connection, _ int, value []byte) {
	// Get the current CSR buffer for the client connection.
	b, ok := p.connRecvBuffers[client]
	if !ok {
		log.Printf("failed to find csr buffer for ble client: %v", client)
		return
	}

	// If all data has been received, handle the CSR data
	if b.recvLen == b.len {
		p.performHandshake(b.Bytes())
		b.Reset()
		return
	}

	log.Printf("received csr data with len = %d: %x", len(value), value)

	// Write the received data to the buffer
	if _, err := b.Write(value); err != nil {
		log.Printf("failed to write csr data to buffer: %v", err)
		return
	}
	b.recvLen += len(value)

	log.Printf("now csr has a total len of %d. want %d", b.Len(), b.len)

	// If all data has been received, handle the CSR data
	if b.recvLen == b.len {
		defer b.Reset()
		log.Printf("csr data: %x", b.Bytes())
		err := p.performHandshake(b.Bytes())
		if err != nil {
			log.Printf("failed to perform handshake: %v", err)
		}
		return
	}
}

// writeCert takes an arbitrary number of bytes in a slice and send them as chucks of 128 bytes to
// the smartphone app.
func (p *Peripheral) writeCert(data []byte) error {
	uintLen := binary.LittleEndian.AppendUint16([]byte{}, uint16(len(data)))
	if _, err := p.lenCertChar.Write(uintLen); err != nil {
		return fmt.Errorf("failed to write cert len data: %v", err)
	}

	// Wait for the smartphone app to be ready to receive the certificate data.
	time.Sleep(250 * time.Millisecond)

	// Send the data in chunks of 128 bytes.
	for i := 0; i < len(data); i += 128 {
		end := i + 128
		if end > len(data) {
			end = len(data)
		}

		if _, err := p.certChar.Write(data[i:end]); err != nil {
			return fmt.Errorf("failed to write cert data: %v", err)
		}

		// Wait for the smartphone app to be ready to receive the next chunk of data.
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}
