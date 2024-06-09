package main

import "fmt"

func CalculateAndPrintStatistics(data []int) {
	average := CalculateAverage(data)
	median := CalculateMedian(data)
	variance := CalculateVariance(data, average)
	stdDeviation := CalculateStandardDeviation(variance)

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", stdDeviation)
}
