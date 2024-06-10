package statcalc

func CalculateVariance(data []int, mean int) int {
	sumOfSquares := 0
	for _, value := range data {
		diff := value - mean
		sumOfSquares += diff * diff
	}
	return sumOfSquares / len(data)
}
