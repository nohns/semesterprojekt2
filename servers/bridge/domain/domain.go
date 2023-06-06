package domain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type domain struct {
	uart uart
}

type uart interface {
	AwaitResponse(cmd int) ([]byte, error)
}

func New(uart uart) *domain {

	return &domain{uart: uart}

}

// Called by bluetooth handler returns the certificate
func (d domain) SignCertificate(rawPubKey []byte) ([]byte, error) {

	//convert csr to x509.CertificateRequest
	/*block, rest := pem.Decode(rawPubKey)
	if rest != nil {
		return nil, fmt.Errorf("could not decode csr pem data: %v", rest)
	}
	if block.Type != "CERTIFICATE REQUEST" {
		return nil, fmt.Errorf("invalid csr pem data type: %s", block.Type)
	}

	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse csr: %v", err)

	}*/

	pubKey, err := x509.ParsePKCS1PublicKey(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %v", err)
	}

	//Load in the intermediate CA's private key and certificate
	caPrivKeyPath, err := filepath.Abs("./cert/ca-key.pem")
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path to ca private key: %v", err)
	}
	caPrivateKey, err := loadPrivateKeyFromFile(caPrivKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not load ca private key: %v", err)
	}

	caPath, err := filepath.Abs("./cert/ca-cert.pem")
	if err != nil {
		return nil, fmt.Errorf("could not get absolute path to ca certificate: %v", err)
	}
	caCert, err := loadCertificateFromFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("could not load ca certificate: %v", err)
	}

	// Create a new certificate template
	template := &x509.Certificate{
		Subject:               pkix.Name{CommonName: "PHONE-CERT"},
		PublicKeyAlgorithm:    x509.RSA,
		SignatureAlgorithm:    x509.UnknownSignatureAlgorithm,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // Valid for 1 year
		SerialNumber:          big.NewInt(1),
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	// Sign the CSR using the CA's private key
	derCert, err := x509.CreateCertificate(rand.Reader, template, caCert, pubKey, caPrivateKey)
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

	block, _ := pem.Decode(pemFile)
	if block == nil {
		return nil, fmt.Errorf("could not decode pem file: %v", err)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("could not cast private key to rsa private key")
	}

	return rsaPrivateKey, nil
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
	block, _ := pem.Decode(pemFile)
	if block == nil {
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
