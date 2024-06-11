package statcalc

import (
	"testing"
)

// TestCalculateVariance tests the CalculateVariance function.
func TestCalculateVariance(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		mean     int
		expected int
	}{
		{
			name:     "Positive numbers with zero variance",
			data:     []int{2, 2, 2, 2},
			mean:     2,
			expected: 0,
		},
		{
			name:     "Positive numbers with non-zero variance",
			data:     []int{1, 2, 3, 4, 5},
			mean:     3,
			expected: 2,
		},
		{
			name:     "Negative numbers with zero variance",
			data:     []int{-3, -3, -3},
			mean:     -3,
			expected: 0,
		},
		{
			name:     "Negative numbers with non-zero variance",
			data:     []int{-5, -1, -3, -2, -4},
			mean:     -3,
			expected: 2,
		},
		{
			name:     "Mixed positive and negative numbers",
			data:     []int{-3, -2, 0, 2, 3},
			mean:     0,
			expected: 5,
		},
		{
			name:     "Single element",
			data:     []int{42},
			mean:     42,
			expected: 0,
		},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateVariance(tt.data, tt.mean)
			if result != tt.expected {
				t.Errorf("CalculateVariance(%v, %d) = %d; want %d", tt.data, tt.mean, result, tt.expected)
			}
		})
	}
}
