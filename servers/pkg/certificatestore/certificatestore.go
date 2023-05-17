package certificatestore

import (
	"database/sql"
	"log"
)

type CertificateStore struct {
	db *sql.DB
}

func New(db *sql.DB) *CertificateStore {
	return &CertificateStore{
		db: db,
	}
}

// TODO
func (cs *CertificateStore) InsertCertificate(id string, certificate []byte, privatekey []byte) error {
	_, err := cs.db.Exec("INSERT INTO certificate (id, certificate, privatekey) VALUES (?, ?, ?)", id, certificate, privatekey)
	if err != nil {
		return err
	}
	return nil
}

// TODO
func (cs *CertificateStore) GetCertificate(id string) ([]byte, []byte, error) {
	//use the id as a string to look into the certificate table and retrieve the certificate and private key
	log.Println("Querying database for certificate and private key")
	result := cs.db.QueryRow("SELECT certificate, privatekey FROM certificate WHERE id = ?", id)
	//return the certificate and private key
	var certificate []byte
	var privatekey []byte
	var err error
	if err = result.Scan(&certificate, &privatekey); err == sql.ErrNoRows {
		return nil, nil, err
	}

	return certificate, privatekey, nil

}
