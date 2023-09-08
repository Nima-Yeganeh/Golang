package main

import (
	"fmt"
)

func fibonacci(limit int) []int {
	sequence := make([]int, 0, limit)
	a, b := 0, 1
	for a <= limit {
		sequence = append(sequence, a)
		a, b = b, a+b
	}
	return sequence
}

func main() {
	limit := 1000 // Change this to your desired limit

	// Generate the Fibonacci sequence up to the specified limit
	fibSeq := fibonacci(limit)

	// Print the Fibonacci numbers
	fmt.Printf("Fibonacci sequence up to %d:\n", limit)
	for _, num := range fibSeq {
		fmt.Println(num)
	}
}
