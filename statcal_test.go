package main

import (
	"math"
	"testing"

	sc "math-skills/statcalc"
)

func TestCalculateAverage(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expected := 3
	result := sc.CalculateAverage(data)
	if result != expected {
		t.Errorf("calculateAverage(%v) = %d; want %d", data, result, expected)
	}
}

// TestCalculateMedian tests the calculateMedian function
func TestCalculateMedian(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expected := 3
	result := sc.CalculateMedian(data)
	if result != expected {
		t.Errorf("calculateMedian(%v) = %d; want %d", data, result, expected)
	}

	data = []float64{1, 2, 3, 4, 5, 6}
	expected = 4
	result = sc.CalculateMedian(data)
	if result != expected {
		t.Errorf("calculateMedian(%v) = %d; want %d", data, result, expected)
	}
}

// TestCalculateVariance tests the calculateVariance function
func TestCalculateVariance(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	mean := sc.CalculateAverage(data)
	expected := 2
	result := sc.CalculateVariance(data, mean)
	if result != expected {
		t.Errorf("calculateVariance(%v, %d) = %d; want %d", data, mean, result, expected)
	}
}

// TestCalculateStandardDeviation tests the calculateStandardDeviation function
func TestCalculateStandardDeviation(t *testing.T) {
	variance := 4.0
	expected := 2
	result := sc.CalculateStandardDeviation(int(math.Round(variance)))
	if result != expected {
		t.Errorf("calculateStandardDeviation(%f) = %d; want %d", variance, result, expected)
	}
}
