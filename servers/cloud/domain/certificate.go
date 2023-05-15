package domain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func createRootCertificate() ([]byte, []byte, error) {
	rootTemplate := x509.Certificate{
		IsCA:         true,
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "Cloud Root CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // valid for 10 years
		BasicConstraintsValid: true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
	}

	rootPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		//format new error
		return nil, nil, fmt.Errorf("failed to generate root private key: %w", err)
	}

	rootCertificateDER, err := x509.CreateCertificate(rand.Reader, &rootTemplate, &rootTemplate, &rootPrivateKey.PublicKey, rootPrivateKey)
	if err != nil {
		//format new error

		return nil, nil, fmt.Errorf("failed to create root certificate: %w", err)
	}

	rootCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rootCertificateDER,
	})

	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rootPrivateKey),
	})

	return rootCertPEM, rootKeyPEM, nil

}

func signCertificate(rootCertPEM, rootKeyPEM, bridgeCSR, bridgeCSRBlock []byte) ([]byte, error) {

	rootCertBlock, rest := pem.Decode(rootCertPEM)
	if len(rest) > 0 {
		return nil, fmt.Errorf("failed to decode root certificate: %w", rest)
	}

	rootCert, err := x509.ParseCertificate(rootCertBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse root certificate: %w", err)
	}

	rootKeyBlock, rest := pem.Decode(rootKeyPEM)
	if len(rest) > 0 {
		return nil, fmt.Errorf("failed to decode root key: %w", rest)
	}

	rootKey, err := x509.ParsePKCS1PrivateKey(rootKeyBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse root key: %w", err)
	}

	bridgeTemplate := x509.Certificate{
		Subject: pkix.Name{
			CommonName: "Bridge",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(2, 0, 0), // valid for 2 years
		KeyUsage:  x509.KeyUsageDigitalSignature,
	}

	bridgeCertificateDER, err := x509.CreateCertificate(rand.Reader, &bridgeTemplate, rootCert, bridgeCSR.PublicKey, rootKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create bridge certificate: %w", err)
	}

	return bridgeCertificateDER, nil
}
