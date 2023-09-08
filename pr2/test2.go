package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Define the number for which you want to calculate the factorial
	num := 5

	// Calculate the factorial
	result := factorial(num)

	// Print the result
	fmt.Printf("The factorial of %d is %d\n", num, result)
}
