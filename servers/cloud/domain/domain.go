package domain

import (
	"crypto/tls"
	"log"

	"github.com/nohns/servers/pkg/certificate"
)

const cloudID = "123"

type CertificateStore interface {
	InsertCertificate(id string, certificate []byte, privatekey []byte) error
	GetCertificate(id string) ([]byte, []byte, error)
}

type domain struct {
	cs CertificateStore
}

func New(certificateStore CertificateStore) *domain {
	return &domain{
		cs: certificateStore,
	}
}

// Unexported method used to initialize the root certificate
func (d domain) InitializeRootCertificate() (*tls.Certificate, error) {

	log.Println("Trying to retrieve root certificate from database")
	rootCertificatePEM, rootPrivateKeyPEM, err := d.cs.GetCertificate(cloudID)
	if err != nil {
		log.Println("Failed to retrieve root certificate from database", err)
		// If no root certificate exists, we generate one
		rootCertificatePEM, rootPrivateKeyPEM, err = certificate.CreateRootCertificate()
		if err != nil {
			//Handle error
			log.Println("Failed to create certificate", err)
		}

		//Store certificate in database
		err = d.cs.InsertCertificate(cloudID, rootCertificatePEM, rootPrivateKeyPEM)
		if err != nil {
			//Handle error
			log.Println("Failed to store certificate in database", err)
		}

	}

	tlsCert, err := certificate.CreateTLSCertFromPEM(rootCertificatePEM, rootPrivateKeyPEM)
	if err != nil {
		log.Println("Failed to create TLS certificate", err)
		return nil, err
	}

	return tlsCert, nil
}

// Unexported mrthod used to sign a certificate
func (d domain) signCertificate(csrPEM []byte) ([]byte, error) {
	//Check database if root certificate exists
	rootCertificatePEM, rootPrivateKeyPEM, err := d.cs.GetCertificate(cloudID)
	if err != nil {
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	//Extract private key from PEM
	tlsCertificate, err := tls.X509KeyPair(rootCertificatePEM, rootPrivateKeyPEM)
	if err != nil {
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	signedCertificatePEM, err := certificate.SignCertificate(rootCertificatePEM, csrPEM, tlsCertificate.PrivateKey)
	if err != nil {
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	return signedCertificatePEM, nil
}

// Associated with the register method in the pairing service
// This method will get hit by the bridge when it wants to pair with the cloud
func (d domain) PairCloud(csrPEM []byte) ([]byte, error) {

	signedCertificatePEM, err := d.signCertificate(csrPEM)
	if err != nil {
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	return signedCertificatePEM, nil
}
