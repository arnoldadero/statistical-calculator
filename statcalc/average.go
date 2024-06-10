package statcalc

func CalculateAverage(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum / len(data)
}
