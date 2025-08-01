package scanner

import (
	"fmt"
	"net"
	"os"
)

// MDNSScan performs an mDNS scan on the local network
func MDNSScan() {
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.251:5353")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		fmt.Printf("Received mDNS packet from %s: %s\n", src, string(buf[:n]))
	}
}
