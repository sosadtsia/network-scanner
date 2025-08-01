package scanner

import (
	"fmt"
	"net"
	"os"
)

// SSDPScan performs an SSDP/UPnP scan on the local network
func SSDPScan() {
	addr, err := net.ResolveUDPAddr("udp", "239.255.255.250:1900")
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

	msg := "M-SEARCH * HTTP/1.1\r\n" +
		"HOST: 239.255.255.250:1900\r\n" +
		"MAN: \"ssdp:discover\"\r\n" +
		"MX: 1\r\n" +
		"ST: ssdp:all\r\n\r\n"

	if _, err := conn.WriteToUDP([]byte(msg), addr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}
		fmt.Printf("Received SSDP packet from %s: %s\n", src, string(buf[:n]))
	}
}
