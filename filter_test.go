package dsp

import (
	"reflect"
	"testing"
)

type test struct {
	A, B, signal, want []float64
}

func TestFiltFilt(t *testing.T) {
	cases := []test{
		{
			[]float64{1, -1.77863177782459, 0.800802646665708},
			[]float64{0.00554271721028068, 0.0110854344205614, 0.00554271721028068},
			[]float64{1, 2, 3, 7, 4, 3, 2, 1, 9, 3, 2, 1},
			[]float64{1.756427537718209, 1.906015981493915, 2.0411967601756045, 2.1566443401702267, 2.248049775459058, 2.3126649846326925, 2.3490173060494457, 2.3561413784140064, 2.333211850462598, 2.280114716164253, 2.1981519700135674, 2.0896901629161326},
		},
	}
	for _, c := range cases {
		got := Filtfilt(c.B, c.A, c.signal)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Filter test failed: got %+v, want %+v", got, c.want)
		}
	}
}
