package dsp

import (
	"math"
)

// Span generates `size` equidistant points spanning [min,max]
// e.g. Span(0, 1, 3) => [0, 0.5, 1]
func Span(min, max float64, size int) []float64 {
	points := make([]float64, size)
	min, max = math.Min(max, min), math.Max(max, min)
	d := max - min
	for i := range points {
		points[i] = min + d*(float64(i)/float64(size-1))
	}
	return points
}

// Piecewise interpolation
//
// Time complexity: O(log(N)), where N is the number of points.
// Space complexity: O(1)
func Interpolate(values []float64, size int) []float64 {
	// Handle special cases
	if len(values) == 0 {
		return values
	}
	if size == 0 {
		return []float64{}
	}

	oldSize := len(values)
	axisX := Span(0, 1, oldSize)
	axisY := values
	valuesNew := make([]float64, size)

	var x, w float64
	var i, j, h int
	for index := 1; index < size-1; index++ {
		x = float64(index) / (float64(size) - 1)
		i, j = 0, oldSize
		for i < j {
			h = int(uint(i+j) >> 1)
			if axisX[h] < x {
				i = h + 1
			} else {
				j = h
			}
		}
		if i < oldSize {
			w = (x - axisX[i-1]) / (axisX[i] - axisX[i-1])
			valuesNew[index] = (1-w)*axisY[i-1] + w*axisY[i]
		} else {
			valuesNew[index] = 0
		}
	}

	// The new values should always have same first value and last value
	valuesNew[0] = values[0]
	valuesNew[size-1] = values[oldSize-1]
	return valuesNew
}
