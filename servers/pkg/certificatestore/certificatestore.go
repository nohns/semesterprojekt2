package certificatestore

import (
	"database/sql"
	"fmt"
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

	result, err := cs.db.Query("SELECT certificate, privatekey FROM certificate WHERE id = ?", id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to select certificate and privatekey based on id: %w", err)
	}
	defer result.Close()
	//return the certificate and private key
	var certificate []byte
	var privatekey []byte
	for result.Next() {
		err := result.Scan(&certificate, &privatekey)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan certificate and privatekey: %w", err)
		}
	}
	return certificate, privatekey, nil

}
