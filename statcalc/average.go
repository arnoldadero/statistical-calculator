package statcalc

import "math"

func CalculateAverage(data []float64) int {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return int(math.Round(sum / float64(len(data))))
}
