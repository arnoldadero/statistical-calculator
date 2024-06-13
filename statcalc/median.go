package statcalc

import (
	"math"
	"sort"
)

func CalculateMedian(data []float64) int {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		return int(math.Round((data[n/2-1] + data[n/2]) / 2))
	}
	return int(math.Round(data[n/2]))
}
