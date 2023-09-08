package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Choose an operation:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		// Read user's choice
		var choice int
		fmt.Print("Enter your choice (1-5): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number (1-5).")
			continue
		}

		// Perform the chosen operation
		switch choice {
		case 1:
			result := performAddition()
			fmt.Printf("Result: %f\n", result)
		case 2:
			result := performSubtraction()
			fmt.Printf("Result: %f\n", result)
		case 3:
			result := performMultiplication()
			fmt.Printf("Result: %f\n", result)
		case 4:
			result, err := performDivision()
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %f\n", result)
			}
		case 5:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select a valid option (1-5).")
		}
	}
}

func performAddition() float64 {
	num1, num2 := readTwoNumbers()
	return num1 + num2
}

func performSubtraction() float64 {
	num1, num2 := readTwoNumbers()
	return num1 - num2
}

func performMultiplication() float64 {
	num1, num2 := readTwoNumbers()
	return num1 * num2
}

func performDivision() (float64, error) {
	num1, num2 := readTwoNumbers()
	if num2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return num1 / num2, nil
}

func readTwoNumbers() (float64, float64) {
	var num1, num2 float64
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	return num1, num2
}
