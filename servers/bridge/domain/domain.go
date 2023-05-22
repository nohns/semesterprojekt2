package domain

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

const filepath = ""

type domain struct {
	uart uart
	cs   certificateStore
}

type uart interface {
	AwaitResponse(ctx context.Context, cmd int) ([]byte, error)
}

type certificateStore interface {
	InsertCertificate(id string, certificate []byte, privatekey []byte) error
	GetCertificate(id string) ([]byte, []byte, error)
}

func New(uart uart, cs certificateStore) *domain {

	return &domain{uart: uart, cs: cs}

}

// Called by bluetooth handler returns the certificate
func (d domain) SignCertificate(csrReq []byte) ([]byte, error) {

	//convert csr to x509.CertificateRequest
	block, rest := pem.Decode(csrReq)
	if rest != nil {
		return nil, fmt.Errorf("could not decode csr pem data: %v", rest)
	}
	if block.Type != "CERTIFICATE REQUEST" {
		return nil, fmt.Errorf("invalid csr pem data type: %s", block.Type)
	}

	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse csr: %v", err)

	}

	//Load in the intermediate CA's private key and certificate
	caPrivateKey, err := loadPrivateKeyFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not load ca private key: %v", err)
	}

	caCert, err := loadCertificateFromFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not load ca certificate: %v", err)
	}

	// Create a new certificate template
	template := &x509.Certificate{
		Subject:               csr.Subject,
		PublicKeyAlgorithm:    csr.PublicKeyAlgorithm,
		PublicKey:             csr.PublicKey,
		SignatureAlgorithm:    csr.SignatureAlgorithm,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // Valid for 1 year
		SerialNumber:          big.NewInt(1),
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	// Sign the CSR using the CA's private key
	derCert, err := x509.CreateCertificate(rand.Reader, template, caCert, csr.PublicKey, caPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("could not sign csr: %v", err)
	}

	// Generate the final certificate in PEM format
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derCert})

	return certPEM, nil
}

func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	pemFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	block, rest := pem.Decode(pemFile)
	if rest != nil {
		return nil, fmt.Errorf("could not decode pem file: %v", err)
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadCertificateFromFile(filename string) (*x509.Certificate, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	defer file.Close()

	pemFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}
	block, rest := pem.Decode(pemFile)
	if rest != nil {
		return nil, fmt.Errorf("could not decode pem file: %v", err)
	}

	return x509.ParseCertificate(block.Bytes)
}

/* func generateCSR(privateKey *rsa.PrivateKey, emailAddress string) []byte {
	subject := pkix.Name{
		CommonName:         "localhost",
		Country:            []string{"DK"},
		Province:           []string{"EU"},
		Locality:           []string{"Copenhagen"},
		Organization:       []string{"Dev"},
		OrganizationalUnit: []string{"Semesterprojekt"},
	}
	template := x509.CertificateRequest{
		Subject:            subject,
		EmailAddresses:     []string{emailAddress},
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &template, privateKey)
	return csrBytes
} */
