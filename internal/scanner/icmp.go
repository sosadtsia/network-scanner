package scanner

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// ICMPScan performs an ICMP scan on a range of IP addresses in the local network
func ICMPScan(network string) {
	fmt.Printf("Starting ICMP scan on network %s...\n", network)
	var wg sync.WaitGroup
	for i := 1; i <= 254; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ip := fmt.Sprintf("%s.%d", network, i)
			conn, err := net.DialTimeout("ip4:icmp", ip, time.Second*2)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error dialing ICMP for %s: %v\n", ip, err)
				return
			}
			defer conn.Close()
			fmt.Printf("ICMP connection established with %s\n", ip)

			// ICMP Echo Request
			msg := []byte{8, 0, 0, 0, 0, 0, 0, 0}
			if _, err := conn.Write(msg); err != nil {
				fmt.Fprintf(os.Stderr, "Error sending ICMP request to %s: %v\n", ip, err)
				return
			}
			fmt.Printf("ICMP Echo Request sent to %s\n", ip)

			// Read ICMP Echo Reply
			buf := make([]byte, 20)
			if _, err := conn.Read(buf); err != nil {
				fmt.Fprintf(os.Stderr, "Error reading ICMP reply from %s: %v\n", ip, err)
				return
			}

			fmt.Printf("Received ICMP reply from %s\n", ip)
		}(i)
	}
	wg.Wait()
	fmt.Println("ICMP scan completed.")
}
