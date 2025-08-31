package main

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"
)

// TestAppRunning verifies that the Fiber app can start and respond to a request.
func TestAppRunning(t *testing.T) {
	app := SetupApp()

	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer ln.Close()

	go func() {
		if err := app.Listener(ln); err != nil && err != http.ErrServerClosed {
			t.Errorf("app.Listener error: %v", err)
		}
	}()

	addr := ln.Addr().(*net.TCPAddr)
	time.Sleep(10 * time.Millisecond)
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/", addr.Port))
	if err != nil {
		t.Fatalf("failed to get: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if err := app.Shutdown(); err != nil {
		t.Fatalf("app.Shutdown error: %v", err)
	}
}
