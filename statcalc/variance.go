package statcalc

import "math"

func CalculateVariance(data []float64, mean int) int {
	meanFloat := float64(mean)
	sumOfSquares := 0.0
	for _, value := range data {
		diff := value - meanFloat
		sumOfSquares += diff * diff
	}
	return int(math.Round(sumOfSquares / float64(len(data))))
}
