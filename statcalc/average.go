package statcalc

func CalculateAverage(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
	// return int(math.Round(sum / float64(len(data))))
}
