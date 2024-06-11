package statcalc

import (
	"math"
	"testing"
)

// TestCalculateStandardDeviation tests the CalculateStandardDeviation function.
func TestCalculateStandardDeviation(t *testing.T) {
	tests := []struct {
		name     string
		variance int
		expected int
	}{
		{
			name:     "Zero variance",
			variance: 0,
			expected: 0,
		},
		{
			name:     "Perfect square variance",
			variance: 4,
			expected: 2,
		},
		{
			name:     "Non-perfect square variance",
			variance: 8,
			expected: int(math.Sqrt(8)),
		},
		{
			name:     "Large perfect square variance",
			variance: 10000,
			expected: 100,
		},
		{
			name:     "Large non-perfect square variance",
			variance: 10001,
			expected: int(math.Sqrt(10001)),
		},
		{
			name:     "Variance of 1",
			variance: 1,
			expected: 1,
		},
		// {
		// 	name:     "Negative variance (invalid case)",
		// 	variance: -4,
		// 	expected: int(math.Sqrt(4)), // Expect the function to handle negative input gracefully
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateStandardDeviation(tt.variance)
			if result != tt.expected {
				t.Errorf("CalculateStandardDeviation(%d) = %d; want %d", tt.variance, result, tt.expected)
			}
		})
	}
}
