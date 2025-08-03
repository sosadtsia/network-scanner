package scanner

import (
	"fmt"
	"os/exec"
)

// ARPScan performs an ARP scan on the local network
func ARPScan() {
	fmt.Println("Executing ARP scan...")
	cmd := exec.Command("arp", "-a")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing arp command: %v\n", err)
		return
	}
	fmt.Println("ARP command executed successfully.")
	fmt.Printf("ARP Scan Results:\n%s\n", string(output))
}
