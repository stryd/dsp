package dsp

import "testing"

type testCase struct {
	points        []Point
	wantSlope     float64
	wantIntercept float64
}

func TestLeastSquares(t *testing.T) {
	cases := []testCase{
		testCase{
			[]Point{
				{1, 6},
				{2, 5},
				{3, 7},
				{4, 10},
			}, 1.4, 3.5,
		},
		testCase{
			[]Point{
				{1, 10},
				{2, 8},
				{3, 7},
				{4, 5},
			}, -1.6, 11.5,
		},
	}
	for _, c := range cases {
		slope, intercept := LeastSquares(&c.points)
		if slope != c.wantSlope || intercept != c.wantIntercept {
			t.Errorf("Least squares calculation is wrong, for slope, want %v, got %v; for intercept, want %v, got %v", c.wantSlope, slope, c.wantIntercept, intercept)
		}
	}

}
