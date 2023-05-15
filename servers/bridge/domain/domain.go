package domain

import "context"

type domain struct {
	uart uart
	cs   certificateStore
}

type uart interface {
	AwaitResponse(ctx context.Context, cmd int) ([]byte, error)
}

type certificateStore interface {
	InsertCertificate(id string, certificate []byte, privatekey []byte) error
	GetCertificate(id string) ([]byte, []byte, error)
}

func New(uart uart, cs certificateStore) *domain {
	return &domain{uart: uart, cs: cs}
}
