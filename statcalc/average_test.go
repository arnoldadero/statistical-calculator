package statcalc

import (
	"testing"
)

// TestCalculateAverage tests the CalculateAverage function.
func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -3,
		},
		{
			name:     "Mixed numbers",
			input:    []int{-1, -2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "Single element",
			input:    []int{10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateAverage(tt.input)
			if result != tt.expected {
				t.Errorf("CalculateAverage(%v) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
