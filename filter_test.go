package dsp

import (
	"math"
	"testing"
)

type test struct {
	freq   float64
	signal []float64
	want   []float64
}

func TestIITFilter(t *testing.T) {
	cases := []test{
		{
			0.05,
			[]float64{1, 2, 3, 7, 4, 3, 2, 1, 9, 3, 2, 1},
			[]float64{1.7248, 2.1277, 2.5071, 2.8313, 3.0805, 3.2544, 3.3651, 3.4237, 3.4339, 3.3999, 3.3385, 3.2723},
		},
	}
	for _, c := range cases {
		got := FilterByIIR(c.signal, c.freq)
		for i := range c.want {
			if math.Abs(c.want[i]-got[i]) > 1e-4 {
				t.Errorf("Filter test failed: got %+v, want %+v", got, c.want)
				break
			}
		}
	}
}
