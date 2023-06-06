package domain

import (
	"context"
	"fmt"
	"testing"
)

func TestTranslator(t *testing.T) {
	// Create a new domain object
	d := domain{}

	// Test case 1: b == ack
	state, err := d.translator(context.Background(), ack, 0, true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != true {
		t.Errorf("unexpected state: got %v, want %v", state, true)
	}

	// Test case 2: b == locked
	state, err = d.translator(context.Background(), locked, 0, true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != true {
		t.Errorf("unexpected state: got %v, want %v", state, true)
	}

	// Test case 3: b == unlocked
	state, err = d.translator(context.Background(), unlocked, 0, true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != false {
		t.Errorf("unexpected state: got %v, want %v", state, false)
	}

	// Test case 4: ctx.Err() != nil
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	state, err = d.translator(ctx, unlocked, 0, true)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
	if state != false {
		t.Errorf("unexpected state: got %v, want %v", state, false)
	}
}

type mockUART struct {
	response []byte
	err      error
}

func (m *mockUART) AwaitResponse(mockCmd int) ([]byte, error) {
	//We need to mock the return of the response based on the commands

	if mockCmd == openLock {
		m.response = []byte{byte(ack)}
		return m.response, m.err
	}

	if mockCmd == closeLock {
		m.response = []byte{byte(ack)}
		return m.response, m.err
	}

	//We only mock the test case for locked
	if mockCmd == lockState {
		m.response = []byte{byte(locked)}
		return m.response, m.err
	}
	return nil, nil
}

func TestGetLock(t *testing.T) {
	// Create a new domain object with a mock UART
	mockUART := &mockUART{response: []byte{byte(locked)}}
	d := domain{uart: mockUART}

	// Test case 1: locked
	state, err := d.GetLock(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != true {
		t.Errorf("unexpected state: got %v, want %v", state, true)
	}

	// Test case 2: unlocked
	/* mockUART.response = []byte{byte(unlocked)}
	state, err = d.GetLock(context.Background())
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != false {
		t.Errorf("unexpected state: got %v, want %v", state, false)
	} */

	// Test case 3: error
	mockUART.err = fmt.Errorf("uart error")
	_, err = d.GetLock(context.Background())
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func TestSetLock(t *testing.T) {
	// Create a new domain object with a mock UART
	mockUART := &mockUART{response: []byte{byte(ack)}}
	d := domain{uart: mockUART}

	// Test case 1: lock
	state, err := d.SetLock(context.Background(), true)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != true {
		t.Errorf("unexpected state: got %v, want %v", state, true)
	}
	if len(mockUART.response) != 1 || mockUART.response[0] != ack {
		t.Errorf("unexpected response: got %v, want %v", mockUART.response, []byte{ack})
	}

	// Test case 2: unlock
	mockUART.response = []byte{byte(ack)}
	state, err = d.SetLock(context.Background(), false)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if state != false {
		t.Errorf("unexpected state: got %v, want %v", state, false)
	}
	if len(mockUART.response) != 1 || mockUART.response[0] != ack {
		t.Errorf("unexpected response: got %v, want %v", mockUART.response, []byte{ack})
	}

	// Test case 3: error
	mockUART.err = fmt.Errorf("uart error")
	_, err = d.SetLock(context.Background(), true)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
