package scanner

import (
	"fmt"
	"net"
	"os"
)

// MDNSScan performs an mDNS scan on the local network
func MDNSScan() {
	fmt.Println("Starting mDNS scan...")
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.251:5353")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving UDP address: %v\n", err)
		return
	}
	fmt.Println("Resolved UDP address for mDNS.")

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening on multicast UDP: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Listening for mDNS packets...")

	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from UDP: %v\n", err)
			return
		}
		fmt.Printf("Received mDNS packet from %s: %s\n", src, string(buf[:n]))
	}
}
