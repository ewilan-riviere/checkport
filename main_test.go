package main

import (
	"context"
	"net"
	"strconv"
	"testing"
	"time"

	psnet "github.com/shirou/gopsutil/v4/net"
)

func TestPortDetection(t *testing.T) {
	// Open a port with the standard ‘net’ package
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Unable to open a port for testing: %v", err)
	}
	defer ln.Close()

	_, portStr, _ := net.SplitHostPort(ln.Addr().String())
	targetPort, _ := strconv.Atoi(portStr)

	t.Logf("Test on dynamic port : %d", targetPort)

	// Use gopsutil via the alias 'psnet'
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	connections, err := psnet.ConnectionsWithContext(ctx, "tcp")
	if err != nil {
		t.Fatalf("Error gopsutil: %v", err)
	}

	found := false
	for _, conn := range connections {
		if conn.Laddr.Port == uint32(targetPort) {
			found = true
			t.Logf("✅ Port %d detected", targetPort)
			break
		}
	}

	if !found {
		t.Errorf("❌ Port %d has not been detected", targetPort)
	}
}
