package dsp

import "testing"

func TestLeastSquares(t *testing.T) {
	signal := []Point{
		{1, 6},
		{2, 5},
		{3, 7},
		{4, 10},
	}
	slope, intercept := LeastSquares(&signal)
	if slope != 1.4 || intercept != 3.5 {
		t.Error("Least squares calculation is wrong")
	}

}
