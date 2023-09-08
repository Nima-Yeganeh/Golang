package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input CSV file
	inputFile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(inputFile)

	// Read all records from the input CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading input CSV:", err)
		return
	}

	// Process the data (e.g., add 10 to the second column)
	for i, record := range records {
		if len(record) >= 2 {
			value, _ := strconv.Atoi(record[1])
			value += 10
			records[i][1] = strconv.Itoa(value)
		}
	}

	// Open the output CSV file for writing
	outputFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(outputFile)

	// Write the modified records to the output CSV
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Println("Error writing to output CSV:", err)
		return
	}

	fmt.Println("Processing complete. Results written to output.csv")
}
