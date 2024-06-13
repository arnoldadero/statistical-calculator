package statcalc

import "math"

func CalculateStandardDeviation(variance float64) int {
	return int(math.Round(math.Sqrt(float64(variance))))
}
