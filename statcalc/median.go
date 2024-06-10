package statcalc

import "sort"

func CalculateMedian(data []int) int {
	sort.Ints(data)
	n := len(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}
