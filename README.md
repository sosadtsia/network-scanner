# network-scanner

Network scanner project

## Overview

This CLI tool allows you to perform network discovery using various protocols:
- ICMP/TCP
- mDNS
- SSDP/UPnP
- ARP

## Installation

Ensure you have Go 1.24 or newer installed. Clone the repository and navigate to the project directory.

```bash
cd projects/network-scanner
```

## Usage

Build the CLI tool:

```bash
go build -o network-scanner ./cmd/network-scanner
```

Run the tool with the desired subcommand:

### ICMP Scan

```bash
./network-scanner icmp <host>
```

### mDNS Scan

```bash
./network-scanner mdns
```

### SSDP/UPnP Scan

```bash
./network-scanner ssdp
```

### ARP Scan

```bash
./network-scanner arp
```

## License

This project is licensed under the MIT License.
