package scanner

import (
	"fmt"
	"os/exec"
)

// ARPScan performs an ARP scan on the local network
func ARPScan() {
	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("ARP Scan Results:\n%s\n", string(output))
}
