package certificate

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

// Function which creates a certificate signing request for the bridge for client tls authentication
func CreateTLSCert(bridgeID string) (*tls.Certificate, error) {
	pk, err := rsa.GenerateKey(rand.Reader, 2048)
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
			Organization: []string{""},
			CommonName:   bridgeID,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, pk.Public(), pk)
	if err != nil {
		return nil, err
	}

	return &tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  pk,
	}, nil

}

func DecodePEM(caPEM, pbPEM []byte) (*pem.Block, *pem.Block, error) {
	cert, rest := pem.Decode(caPEM)
	if len(rest) > 0 {
		return nil, nil, fmt.Errorf("failed to decode certificate in PEM format: %x", rest)
	}

	pb, rest := pem.Decode(pbPEM)
	if len(rest) > 0 {
		return nil, nil, fmt.Errorf("failed to decode root key: %x", rest)
	}

	return cert, pb, nil
}

func EncodePEM(csr *tls.Certificate) ([]byte, []byte, error) {
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
}

func CreateTLSCertFromBytes(cer, pk *pem.Block) (*tls.Certificate, error) {
	cert, err := x509.ParseCertificate(cer.Bytes)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(pk.Bytes)
	if err != nil {
		return nil, err
	}

	return &tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
	}, nil
}
