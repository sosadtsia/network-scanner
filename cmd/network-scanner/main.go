package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/svosadtsia/network-scanner/internal/scanner"
)

func main() {
	// Define subcommands
	icmpCmd := flag.NewFlagSet("icmp", flag.ExitOnError)
	mdnsCmd := flag.NewFlagSet("mdns", flag.ExitOnError)
	ssdpCmd := flag.NewFlagSet("ssdp", flag.ExitOnError)
	arpCmd := flag.NewFlagSet("arp", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'icmp', 'mdns', 'ssdp', or 'arp' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "icmp":
		icmpCmd.Parse(os.Args[2:])
		if icmpCmd.NArg() < 1 {
			fmt.Println("Please provide a host to scan.")
			os.Exit(1)
		}
		host := icmpCmd.Arg(0)
		fmt.Printf("Running ICMP scan on %s...\n", host)
		scanner.ICMPScan(host)
	case "mdns":
		mdnsCmd.Parse(os.Args[2:])
		fmt.Println("Running mDNS scan...")
		// Call mDNS scan function
	case "ssdp":
		ssdpCmd.Parse(os.Args[2:])
		fmt.Println("Running SSDP/UPnP scan...")
		// Call SSDP/UPnP scan function
	case "arp":
		arpCmd.Parse(os.Args[2:])
		fmt.Println("Running ARP scan...")
		// Call ARP scan function
	default:
		fmt.Println("expected 'icmp', 'mdns', 'ssdp', or 'arp' subcommands")
		os.Exit(1)
	}
}
