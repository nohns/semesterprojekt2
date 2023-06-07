package uart

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.bug.st/serial"
)

func TestUart_Write(t *testing.T) {
	u := New() // Create a new Uart instance for testing

	var buf bytes.Buffer
	u.port = &mockPort{WriteFunc: func(p []byte) (n int, err error) {
		return buf.Write(p)
	}}

	data := []byte("test")
	err := u.Write(data)
	if err != nil {
		t.Errorf("Write failed: %s", err)
	}

	expected := append(data, '\x00')
	if !bytes.Equal(buf.Bytes(), expected) {
		t.Errorf("Unexpected write data. Got: %v, Expected: %v", buf.Bytes(), expected)
	}
}

func TestUartRead(t *testing.T) {
	// Create a mock port
	mock := &mockPort{
		ReadFunc: func(p []byte) (n int, err error) {
			// Simulate reading data from the port
			copy(p, []byte("Hello, World!\x00"))
			return len(p), nil
		},
	}

	// Create a new Uart instance with the mock port
	uart := &Uart{
		port: mock,
	}

	// Call the Read function
	data, err := uart.Read()

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, []byte("Hello, World!\x00"), data)
}

type mockPort struct {
	WriteFunc              func(p []byte) (n int, err error)
	ReadFunc               func(p []byte) (n int, err error)
	CloseFunc              func() error
	BreakFunc              func(duration time.Duration) error
	ResetInputBufferFunc   func() error
	ResetOutputBufferFunc  func() error
	SetDTRFunc             func(dtr bool) error
	SetRTSFunc             func(rts bool) error
	GetModemStatusBitsFunc func() (*serial.ModemStatusBits, error)
	SetReadTimeoutFunc     func(t time.Duration) error
	SetModeFunc            func(mode *serial.Mode) error
}

func (m *mockPort) Write(p []byte) (n int, err error) {
	if m.WriteFunc != nil {
		return m.WriteFunc(p)
	}
	return 0, nil // Default behavior for mock port
}

func (m *mockPort) Read(p []byte) (n int, err error) {
	if m.ReadFunc != nil {
		return m.ReadFunc(p)
	}
	return 0, nil // Default behavior for mock port
}

func (m *mockPort) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) Break(duration time.Duration) error {
	if m.BreakFunc != nil {
		return m.BreakFunc(duration)
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) ResetInputBuffer() error {
	if m.ResetInputBufferFunc != nil {
		return m.ResetInputBufferFunc()
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) ResetOutputBuffer() error {
	if m.ResetOutputBufferFunc != nil {
		return m.ResetOutputBufferFunc()
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) SetDTR(dtr bool) error {
	if m.SetDTRFunc != nil {
		return m.SetDTRFunc(dtr)
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) SetRTS(rts bool) error {
	if m.SetRTSFunc != nil {
		return m.SetRTSFunc(rts)
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	if m.GetModemStatusBitsFunc != nil {
		return m.GetModemStatusBitsFunc()
	}
	return nil, nil // Default behavior for mock port
}

func (m *mockPort) SetReadTimeout(t time.Duration) error {
	if m.SetReadTimeoutFunc != nil {
		return m.SetReadTimeoutFunc(t)
	}
	return nil // Default behavior for mock port
}

func (m *mockPort) SetMode(mode *serial.Mode) error {
	if m.SetModeFunc != nil {
		return m.SetModeFunc(mode)
	}
	return nil // Default behavior for mock port
}
