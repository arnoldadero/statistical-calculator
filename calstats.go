package main

import (
	"fmt"

	sc "math-skills/statcalc"
) 

func CalculateAndPrintStatistics(data []int) {
	average := sc.CalculateAverage(data)
	median := sc.CalculateMedian(data)
	variance := sc.CalculateVariance(data, average)
	stdDeviation := sc.CalculateStandardDeviation(variance)

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", stdDeviation)
}
