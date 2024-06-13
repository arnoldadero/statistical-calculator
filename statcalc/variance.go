package statcalc

func CalculateVariance(data []float64, mean float64) float64 {
	// meanFloat := float64(mean)
	sumOfSquares := 0.0
	for _, value := range data {
		diff := value - mean
		sumOfSquares += diff * diff
	}
	return sumOfSquares / float64(len(data))
	// return int(math.Round(sumOfSquares / float64(len(data))))
}
