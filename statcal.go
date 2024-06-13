package main

import (
	"fmt"
	"math"

	sc "math-skills/statcalc"
)

func CalculateAndPrintStatistics(data []float64) {
	average := sc.CalculateAverage(data)
	median := sc.CalculateMedian(data)
	variance := sc.CalculateVariance(data, average)
	stdDeviation := sc.CalculateStandardDeviation(variance)

	// rounding average and variance then change to int
	averageInt := int(math.Round(average))
	varianceInt := int(math.Round(variance))

	fmt.Printf("Average: %d\n", averageInt)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", varianceInt)
	fmt.Printf("Standard Deviation: %d\n", stdDeviation)
}
