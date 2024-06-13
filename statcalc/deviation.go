package statcalc

import "math"

func CalculateStandardDeviation(variance int) int {
	return int(math.Round(math.Sqrt(float64(variance))))
}
