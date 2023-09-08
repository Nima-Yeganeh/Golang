package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Get the operating system and architecture
	os := runtime.GOOS
	arch := runtime.GOARCH

	// Get the Go runtime version
	goVersion := runtime.Version()

	fmt.Printf("Operating System: %s\n", os)
	fmt.Printf("Architecture: %s\n", arch)
	fmt.Printf("Go Version: %s\n", goVersion)
}
