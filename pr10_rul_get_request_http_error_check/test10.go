package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.yahoo.com"

	// Send an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// Check if the status code is 200 (OK)
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Website %s is reachable with status code 200 OK.\n", url)
	} else {
		fmt.Printf("Website %s returned status code %d.\n", url, resp.StatusCode)
	}
}
