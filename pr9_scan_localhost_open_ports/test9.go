package main

import (
	"fmt"
	"net"
	"time"
)

func checkPort(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	host := "localhost"
	startPort := 1
	endPort := 1024 // You can specify your own range of ports

	timeout := 2 * time.Second // Adjust the timeout as needed

	fmt.Printf("Scanning ports on %s...\n", host)

	for port := startPort; port <= endPort; port++ {
		if checkPort(host, port, timeout) {
			fmt.Printf("Port %d is open\n", port)
		}
	}
}
