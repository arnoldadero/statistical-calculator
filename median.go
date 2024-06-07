func calculateMedian(data []int) int{
	sort.Int(data)
	n := len(data)
	if n%2 == 0 {
		return(data[n/2-1] + data[n/2])/2
	}
	return data[n/2]
}