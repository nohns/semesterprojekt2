package certificate

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/nohns/servers/bridge"
)

func ParseCSR(pemdata []byte) (*bridge.PhoneSigningRequest, error) {
	block, err := pem.Decode(pemdata)
	if err != nil {
		return nil, fmt.Errorf("could not decode csr pem data: %v", err)
	}
	if block.Type != "CERTIFICATE REQUEST" {
		return nil, fmt.Errorf("invalid csr pem data type: %s", block.Type)
	}

	// Parse CSR
	csr, cerr := x509.ParseCertificateRequest([]byte{})
	if cerr != nil {
		return nil, fmt.Errorf("could not parse csr: %v", err)
	}

	return &bridge.PhoneSigningRequest{
		DeviceID: csr.Subject.CommonName,
		CSR:      csr,
	}, nil
}
