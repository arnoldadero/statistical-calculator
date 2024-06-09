package main

import (
	"log"
	"os"
)

// main function is the entry point of the program
func main() {
	// Ensure a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the data file.")
	}
	filePath := os.Args[1]

	// Read data from the provided file path
	data, err := readDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate and print the statistics
	CalculateAndPrintStatistics(data)
}
