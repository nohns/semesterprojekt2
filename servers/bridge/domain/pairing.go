package domain

import (
	"context"
	"crypto/tls"
	"log"

	pairingv1 "github.com/nohns/proto/pairing/v1"
	"github.com/nohns/servers/bridge/certificate"
	"github.com/nohns/servers/bridge/server"
)

const bridgeID = "123"

func (d domain) Register() (string, error) {
	return "", nil
}

func (d domain) PairCloud(addr string) (*tls.Certificate, error) {
	// First we check the database if a root certificate exists
	cerPEM, pkPEM, err := d.cs.GetCertificate(bridgeID)
	if err != nil {
		// If no root certificate exists, we generate one
		cer, err := certificate.CreateTLSCert(bridgeID)
		if err != nil {
			//Handle error
			log.Println("Failed to create certificate", err)
		}
		//
		client, conn := server.RunWithTLSAuth(addr, cer)

		//Generate new CSR and send it to the cloud for signing
		csr, err := certificate.CreateTLSCert(bridgeID)
		if err != nil {
			//Handle error
			log.Println("Failed to create CSR", err)
		}
		//Encode the CSR to PEM
		csrPEM, pkPEM, err := certificate.EncodePEM(csr)
		if err != nil {
			//Handle error
			log.Println("Failed to encode CSR to PEM", err)
		}

		//Send CSR to cloud
		signedCertifacteData, err := client.Register(context.Background(), &pairingv1.RegisterRequest{
			Csr: csrPEM,
		})
		if err != nil {
			//Handle error
			log.Println("Failed to send CSR to cloud", err)
		}

		//Split the signed certificate into certificate and id
		id := signedCertifacteData.BridgeId
		signedCertificate := signedCertifacteData.Cert

		//Decode PEM encoded certificate
		cert, pk, err := certificate.DecodePEM(signedCertificate, pkPEM)
		if err != nil {
			//Handle error
			log.Println("Failed to decode PEM", err)
		}

		//Store certificate in database
		err = d.cs.InsertCertificate(id, cert.Bytes, pk.Bytes)
		if err != nil {
			//Handle error
			log.Println("Failed to store certificate in database", err)
		}

		//Close connection
		conn.Close()

		//Create TLS certificate
		tlsCert, err := certificate.CreateTLSCertFromBytes(cert, pk)
		if err != nil {
			//Handle error
			log.Println("Failed to create TLS certificate", err)
		}

		return tlsCert, nil

	}
	//Decode PEM encoded certificate and private key
	cer, pk, err := certificate.DecodePEM(cerPEM, pkPEM)
	if err != nil {
		//Handle error
		log.Println("Failed to decode PEM", err)
	}

	//Create TLS certificate
	tlsCert, err := certificate.CreateTLSCertFromBytes(cer, pk)
	if err != nil {
		//Handle error
		log.Println("Failed to create TLS certificate", err)
	}

	return tlsCert, nil

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
