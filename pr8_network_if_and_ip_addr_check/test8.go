package main

import (
	"fmt"
	"net"
)

func main() {
	// Get a list of network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Loop through the network interfaces
	for _, iface := range interfaces {
		fmt.Printf("Interface Name: %s\n", iface.Name)

		// Get the interface's IP addresses
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Loop through the IP addresses associated with the interface
		for _, addr := range addrs {
			fmt.Printf("  IP Address: %s\n", addr)
		}

		fmt.Println()
	}
}
