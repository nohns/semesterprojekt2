package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"testing"
)

func LoadCertificates() error {
	// Read in the CA certificate
	caCert, err := ioutil.ReadFile("../../../cert/ca-cert.pem")
	if err != nil {
		return fmt.Errorf("failed to read CA certificate: %v", err)
	}

	// Create a new certificate pool and add the CA certificate to it
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return fmt.Errorf("failed to append CA certificate to cert pool")
	}

	// Load the client certificate
	clientCert, err := tls.LoadX509KeyPair("../../../cert/client-cert.pem", "../../../cert/client-key.pem")
	if err != nil {
		return fmt.Errorf("failed to load client certificate: %v", err)
	}

	// Create a TLS configuration with the client certificate and certificate pool
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	// Dial the server with the TLS configuration
	conn, err := tls.Dial("tcp", "localhost:1234", config)
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}

	// Close the connection
	defer conn.Close()

	// Check that the connection is valid
	if err := conn.Handshake(); err != nil {
		return fmt.Errorf("failed to handshake with server: %v", err)
	}

	return nil
}

func TestLoadCertificates(t *testing.T) {
	if err := LoadCertificates(); err != nil {
		fmt.Printf("failed to load certificates: %v\n", err)
	}
}
