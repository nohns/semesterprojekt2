package domain

import (
	"context"
	"crypto/tls"
	"log"

	pairingv1 "github.com/nohns/proto/pairing/v1"
	"github.com/nohns/servers/bridge/server"
	"github.com/nohns/servers/pkg/certificate"
)

const bridgeID = "123"

func (d domain) Register() (string, error) {
	return "", nil
}

func (d domain) PairCloud(addr string) (*tls.Certificate, error) {
	// First we check the database if a root certificate exists
	certPEM, pkPEM, err := d.cs.GetCertificate(bridgeID)
	if err != nil {
		// If no root certificate exists, we generate a self signed one with limited validity
		cer, err := certificate.CreateTempTLSCert(bridgeID)
		if err != nil {
			log.Println("Failed to create certificate", err)
			return nil, err
		}
		//
		client, conn := server.RunWithTLSAuth(addr, cer)

		//Generate new CSR and send it to the cloud for signing
		csrTemplatePEM, csrPrivateKeyPEM, err := certificate.CreateCsrRequest()
		if err != nil {
			log.Println("Failed to create CSR", err)
			return nil, err
		}

		//Extract the public key from the CSR template
		csrPublicKeyPEM, err := certificate.ExtractPublicKeyPEM(csrPrivateKeyPEM)
		if err != nil {
			log.Println("Failed to extract public key from CSR template", err)
			return nil, err
		}

		log.Println(string(csrTemplatePEM))
		log.Println(string(csrPrivateKeyPEM))

		//Send CSR to cloud
		signedCertifacteData, err := client.Register(context.Background(), &pairingv1.RegisterRequest{
			Csr:       csrTemplatePEM,
			PublicKey: csrPublicKeyPEM,
		})
		if err != nil {
			log.Println("Failed to send CSR to cloud", err)
			return nil, err
		}
		//THIS IS WHERE THE ERROR IS HAPPENING

		//Split the signed certificate into certificate and id
		id := signedCertifacteData.BridgeId
		signedCertificate := signedCertifacteData.Cert

		//Store certificate in database
		err = d.cs.InsertCertificate(id, signedCertificate, csrPrivateKeyPEM)
		if err != nil {
			//Handle error
			log.Println("Failed to store certificate in database", err)
			return nil, err
		}

		//Close connection
		log.Println("Closing connection")
		conn.Close()

		//Create TLS certificate
		tlsCertificate, err := certificate.CreateTLSCertFromPEM(signedCertificate, csrPrivateKeyPEM)
		if err != nil {
			//Handle error
			log.Println("Failed to create TLS certificate", err)
		}

		return tlsCertificate, nil

	}
	//Create TLS certificate
	tlsCertificate, err := certificate.CreateTLSCertFromPEM(certPEM, pkPEM)
	if err != nil {
		//Handle error
		log.Println("Failed to create TLS certificate", err)
	}

	return tlsCertificate, nil

}

func (d domain) PairDevice() {

}

//Verify root certificate
//If root certificate is valid, load it into memory
//If root certificate is invalid, delete it from database and file

//The flow is as follows:
//We check if we have a certificate in the database
//We instantly return because that means we are already paired

//If we do not have a certificate in the database, we generate one
//And we use the certificate to start a client connection to the cloud
//A secure connection is established between the bridge and the cloud
//The bridge sends a CSR to the cloud
//The cloud signs the CSR and sends the signed certificate back to the bridge
//The bridge stores the signed certificate in the database
//The bridge uses the signed certificate to start a server connection
//The bridge is now paired with the cloud
