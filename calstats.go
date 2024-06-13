package main

import (
	"fmt"

	sc "math-skills/statcalc"
) 

func CalculateAndPrintStatistics(data []float64) {
	average := sc.CalculateAverage(data)
	median := sc.CalculateMedian(data)
	variance := sc.CalculateVariance(data, average)
	stdDeviation := sc.CalculateStandardDeviation(variance)

	fmt.Printf("Average: %f\n", average)
	fmt.Printf("Median: %f\n", median)
	fmt.Printf("Variance: %f\n", variance)
	fmt.Printf("Standard Deviation: %f\n", stdDeviation)
}
