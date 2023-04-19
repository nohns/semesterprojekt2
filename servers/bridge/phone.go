package bridge

import "crypto/x509"

// Phone is an aggregate root for representing connected phones to the bridge.
type Phone struct {
	ID       string // UUID of phone
	DeviceID string // Device ID of phone
	Cert     []byte // CA self-signed certificate for phone
}

type PhoneSigningRequest struct {
	DeviceID string                   // Device ID of phone
	CSR      *x509.CertificateRequest // Certificate signing request
}
