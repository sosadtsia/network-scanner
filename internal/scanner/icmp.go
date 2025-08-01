package scanner

import (
	"fmt"
	"net"
	"os"
	"time"
)

// ICMPScan performs an ICMP scan on the specified host
func ICMPScan(host string) {
	conn, err := net.DialTimeout("ip4:icmp", host, time.Second*2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer conn.Close()

	// ICMP Echo Request
	msg := []byte{8, 0, 0, 0, 0, 0, 0, 0}
	if _, err := conn.Write(msg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	// Read ICMP Echo Reply
	buf := make([]byte, 20)
	if _, err := conn.Read(buf); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	fmt.Printf("Received ICMP reply from %s\n", host)
}
