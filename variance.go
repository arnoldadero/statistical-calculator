func calculateVariance(data []int, mean int) int {
	sumOfSquares := 0
	 for _, value := range data {
		diff := value - mean
		sumOfSquares += diff * diff
	 }
	 return somOfSquares / len(data)
}