package certificate

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"time"
)

// Function that creates a root certificate for the cloud. It returns the certificate and the private key encoded in PEM format.
func CreateRootCertificate() ([]byte, []byte, error) {

	//Create new serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	//Create root certificate template
	rootTemplate := x509.Certificate{
		IsCA:         true,
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: "Cloud Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // valid for 10 years
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	//generate root private /public keypair
	rootPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		//format new error
		return nil, nil, fmt.Errorf("failed to generate root private key: %w", err)
	}

	//create root certificate
	rootCertificate, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootPrivateKey.PublicKey, rootPrivateKey)
	if err != nil {
		//format new error
		return nil, nil, fmt.Errorf("failed to create root certificate: %w", err)
	}

	rootCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rootCertificate,
	})

	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rootPrivateKey),
	})

	if rootCertPEM == nil || rootKeyPEM == nil {
		log.Println("Failed to encode root certificate or private key")
		return nil, nil, fmt.Errorf("failed to encode root certificate or private key")
	}

	return rootCertPEM, rootKeyPEM, nil

}

// Function that creates a CSR. It returns the CSR template and the private key encoded in PEM format.
func CreateCsrRequest() ([]byte, []byte, error) {
	//Create new serial number
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, nil, err
	}

	//Create ca certificate template
	csrTemplate := x509.Certificate{
		IsCA:         false,
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: "Bridge CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // valid for 10 years
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	//generate root private /public keypair
	csrPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		//format new error
		return nil, nil, fmt.Errorf("failed to generate root private key: %w", err)
	}

	// Encode intermediary certificate template in PEM format
	csrTemplatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: csrTemplate.Raw,
	})

	// Encode intermediary private key in PEM format
	csrKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(csrPrivateKey),
	})

	return csrTemplatePEM, csrKeyPEM, nil

}

// Function which creates a certificate signing request for the bridge for client tls authentication
func CreateTempTLSCert(id string) (*tls.Certificate, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Bridge"},
			CommonName:   id,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(1 * time.Hour), // valid for 1 hour
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, privatekey.Public(), privatekey)
	if err != nil {
		return nil, err
	}

	return &tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  privatekey,
	}, nil
}

// Function which takes a root certificate and a template for an intermediary certificate and signs the intermediary certificate with the root certificate
func SignCertificate(rootCertPEM []byte, csrPEM []byte, rootPrivateKey crypto.PrivateKey) ([]byte, error) {
	// Parse the root certificate
	rootCertBlock, _ := pem.Decode(rootCertPEM)
	if rootCertBlock == nil || rootCertBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	rootCert, err := x509.ParseCertificate(rootCertBlock.Bytes)
	if err != nil {
		return nil, err
	}

	// parse the csr template
	csrTemplate, err := x509.ParseCertificate(csrPEM)
	if err != nil {
		return nil, err
	}

	//Perform changes to the template to make it an intermediary certificate
	csrTemplate.Issuer = rootCert.Subject

	validFrom := time.Now()
	validTo := validFrom.Add(365 * 24 * time.Hour)
	csrTemplate.NotBefore = validFrom
	csrTemplate.NotAfter = validTo

	// Create the intermediary certificate by signing it with the root certificate and key
	intermediaryCert, err := x509.CreateCertificate(rand.Reader, csrTemplate, rootCert, csrTemplate.PublicKey, rootPrivateKey)
	if err != nil {
		return nil, err
	}

	// Encode the intermediary certificate in PEM format
	intermediaryCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: intermediaryCert,
	})

	return intermediaryCertPEM, nil
}

// Unsure if this function is needed tbh
func ExtractPublicKeyPEM(privateKeyPEM []byte) ([]byte, error) {
	// Decode the PEM-encoded private key
	block, _ := pem.Decode(privateKeyPEM)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	// Parse the RSA private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
	}

	// Marshal the public key to DER format
	publicKeyDER, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	// Encode the public key to PEM format
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDER,
	})

	return publicKeyPEM, nil
}

func CreateTLSCertFromPEM(certificatePEM, privateKeyPEM []byte) (*tls.Certificate, error) {

	log.Println("1: ", string(privateKeyPEM))
	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	log.Println("2: ", string(certificatePEM))
	certificateBlock, _ := pem.Decode(certificatePEM)
	if certificateBlock == nil || certificateBlock.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("failed to decode PEM block containing certificate")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
	}

	certificate, err := x509.ParseCertificate(certificateBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	tlsCertificate := &tls.Certificate{
		Certificate: [][]byte{certificate.Raw},
		PrivateKey:  privateKey,
	}

	return tlsCertificate, nil
}

/* func EncodePEM(csr *tls.Certificate) ([]byte, []byte, error) {
	pemCert := new(bytes.Buffer)
	for _, certBytes := range csr.Certificate {
		err := pem.Encode(pemCert, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
		if err != nil {
			return nil, nil, err
		}
	}

	pemKey := new(bytes.Buffer)
	err := pem.Encode(pemKey, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(csr.PrivateKey.(*rsa.PrivateKey))})
	if err != nil {
		return nil, nil, err
	}

	return pemCert.Bytes(), pemKey.Bytes(), nil
} */
