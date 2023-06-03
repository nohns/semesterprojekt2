package server

import (
	"context"
	"testing"

	lockv1 "github.com/nohns/proto/lock/v1"
)

func TestGetLockState(t *testing.T) {
	// Create a mock domain object
	mockDomain := &mockDomain{
		getLockFunc: func(ctx context.Context) (bool, error) {
			return true, nil
		},
	}

	// Create a new server object with the mock domain
	server := &server{domain: mockDomain}

	// Create a new request object
	req := &lockv1.GetLockStateRequest{Id: "test"}

	// Call the GetLockState function
	resp, err := server.GetLockState(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that the response is correct
	if resp.Engaged != true {
		t.Errorf("unexpected response: %v", resp)
	}
}

func TestSetLockState(t *testing.T) {
	// Create a mock domain object
	mockDomain := &mockDomain{
		setLockFunc: func(ctx context.Context, state bool) (bool, error) {
			return state, nil
		},
	}

	// Create a new server object with the mock domain
	server := &server{domain: mockDomain}

	// Create a new request object
	req := &lockv1.SetLockStateRequest{Id: "test", Engaged: true}

	// Call the SetLockState function
	resp, err := server.SetLockState(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that the response is correct
	if resp.Engaged != true {
		t.Errorf("unexpected response: %v", resp)
	}
}

// Define a mock domain object for testing
type mockDomain struct {
	getLockFunc func(ctx context.Context) (bool, error)
	setLockFunc func(ctx context.Context, state bool) (bool, error)
}

func (m *mockDomain) GetLock(ctx context.Context) (bool, error) {
	return m.getLockFunc(ctx)
}

func (m *mockDomain) SetLock(ctx context.Context, state bool) (bool, error) {
	return m.setLockFunc(ctx, state)
}
