package domain

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
