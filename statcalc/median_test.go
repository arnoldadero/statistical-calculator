package statcalc

import (
	"testing"
)

// TestCalculateMedian tests the CalculateMedian function.
func TestCalculateMedian(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		expected int
	}{
		{
			name:     "Odd length with positive numbers",
			data:     []int{3, 1, 2},
			expected: 2,
		},
		{
			name:     "Even length with positive numbers",
			data:     []int{4, 1, 3, 2},
			expected: 2,
		},
		{
			name:     "Odd length with negative numbers",
			data:     []int{-1, -3, -2},
			expected: -2,
		},
		{
			name:     "Even length with negative numbers",
			data:     []int{-4, -1, -3, -2},
			expected: -2,
		},
		{
			name:     "Mixed positive and negative numbers",
			data:     []int{3, -1, 2, -2, 1},
			expected: 1,
		},
		{
			name:     "Single element",
			data:     []int{1},
			expected: 1,
		},
		// {
		// 	name:     "Empty slice",
		// 	data:     []int{},
		// 	expected: 0, // Convention: return 0 for empty input
		// },
		{
			name:     "Large odd length",
			data:     generateSequentialSlice(10001), // Generates [0, 1, 2, ..., 10000]
			expected: 5000,
		},
		{
			name:     "Large even length",
			data:     generateSequentialSlice(10000), // Generates [0, 1, 2, ..., 9999]
			expected: 4999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateMedian(tt.data)
			if result != tt.expected {
				t.Errorf("CalculateMedian(%v) = %d; want %d", tt.data, result, tt.expected)
			}
		})
	}
}

// generateSequentialSlice generates a slice of sequential integers from 0 to n-1.
func generateSequentialSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}
