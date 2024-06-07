func CalculateStatistics(data[]int){
	average := calculateAverage(data)
    median  := calculateMedian(data)
	variance := calculateVariance(data, average)
	stdDeviation := calculateStandardDeviation(variance)

	fmt.Printf("Average: %d\n", average)
	fmt,Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", stdDeviation)
}