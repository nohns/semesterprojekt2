package domain

import "log"

const cloudID = "123"

type CertificateStore interface {
	InsertCertificate(id string, certificate []byte, privatekey []byte) error
	GetCertificate(id string) ([]byte, []byte, error)
}

type domain struct {
	certificateStore CertificateStore
}

func New(certificateStore CertificateStore) *domain {
	return &domain{
		certificateStore: certificateStore,
	}
}

func (d domain) InitializeRootCertificate() {

	// Check database if root certificate exists
	rootcertificate, privKey, err := d.certificateStore.GetCertificate(cloudID)
	if err != nil {
		//Create root certificate
		rootcertificate, privKey, err = createRootCertificate()
		if err != nil {
			log.Fatalln("Gracefully exiting due to error: ", err)
		}
		//Save root certificate to database
		err = d.certificateStore.InsertCertificate(cloudID, rootcertificate, privKey)
		if err != nil {
			//Handle error
			log.Fatalln("Gracefully exiting due to error: ", err)
		}
	}
	//Load root certificate from database
	//Load root certificate from file
	//Verify root certificate
	//If root certificate is valid, load it into memory
	//If root certificate is invalid, delete it from database and file

}

func (d domain) SignCertificate(bridgeCSR, bridgeCSRBlock []byte) ([]byte, error) {
	//Check database if root certificate exists
	rootCertPEM, rootKeyPEM, err := d.certificateStore.GetCertificate(cloudID)
	if err != nil {
		//Handle error
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	signedCertificate, err := signCertificate(rootCertPEM, rootKeyPEM, bridgeCSR, bridgeCSRBlock)
	if err != nil {
		//Handle error
		log.Fatalln("Gracefully exiting due to error: ", err)
	}

	return signedCertificate, nil
}
