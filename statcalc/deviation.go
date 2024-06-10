package statcalc

import "math"

func CalculateStandardDeviation(variance int) int {
	return int(math.Sqrt(float64(variance)))
}
