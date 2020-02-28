package dsp

import (
	"reflect"
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
			[]float64{1.7248337946203223, 2.1277404138921936, 2.5071525898144595, 2.8313057595849633, 3.080590374397801, 3.2544511247177215, 3.3651080012955776, 3.423796407979325, 3.433937135070785, 3.3999620059276547, 3.3385096171139304, 3.2723945635692946},
		},
	}
	for _, c := range cases {
		got := FilterByIIR(c.signal, c.freq)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Filter test failed: got %+v, want %+v", got, c.want)
		}
	}
}
