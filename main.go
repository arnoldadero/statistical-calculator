package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	// Ensure that the filename argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}
	filename := os.Args[1]

	// Read numbers from the file
	numbers, err := readNumbersFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading numbers from file: %v\n", err)
		return
	}
	// Calculate and print each statistical mearsure
	average := calculateAverage(numbers)
	median := calculateMedian(numbers)
	variance := calculateVariance(numbers, average)
	standardDeviation := math.Sqrt(variance)

	// Print the results, rounded to the nearest integer
	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(standardDeviation)))
}
