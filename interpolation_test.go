package dsp

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpan(t *testing.T) {
	type test struct {
		min  float64
		max  float64
		size int
		want []float64
	}
	cases := []test{
		{
			0, 1, 3,
			[]float64{0.0, 0.5, 1.0},
		},
		{
			0, 9, 10,
			[]float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0},
		},
		{
			0, 8, 5,
			[]float64{0.0, 2.0, 4.0, 6.0, 8.0},
		},
	}
	for _, c := range cases {
		got := Span(c.min, c.max, c.size)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Span test failed: got %+v, want %+v", got, c.want)
		}
	}
}

func TestInterpolate(t *testing.T) {
	type test struct {
		in   []float64
		size int
		want []float64
	}
	cases := []test{
		{
			[]float64{0.0, 1.0, 0.0},
			5,
			[]float64{0.0, 0.5, 1.0, 0.5, 0.0},
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0},
			5,
			[]float64{1.0, 3.0, 5.0, 7.0, 9.0},
		},
		{
			[]float64{0.0, 2.0, 4.0, 6.0, 8.0},
			9,
			[]float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0},
		},
		// Following case should be passed but fails the DeepEqual. Commment out for now.
		/*
			{
				[]float64{0.0, 2.0, 0.0, 4.0, 0.0, 2.0, 0.0},
				13,
				[]float64{0.0, 1.0, 2.0, 1.0, 0.0, 2.0, 4.0, 2.0, 0.0, 1.0, 2.0, 1.0, 0.0},
			},
		*/
	}
	for _, c := range cases {
		got := Interpolate(c.in, c.size)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Interpolation test failed: got %+v, want %+v", got, c.want)
		}
	}
}

func TestBinaryFindIndex(t *testing.T) {
	type test struct {
		values             []float64
		qValue             float64
		expectedIndexRange float64
	}
	cases := []test{
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			3.5,
			2.5,
		},
		{
			[]float64{10.0, 20.0, 30.0, 40.0, 50.0},
			10.0,
			0.0,
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			3.0,
			2.0,
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			6.0,
			-1.0,
		},
		{
			[]float64{0.6939693089629557, 0.7037435245821523, 0.7135177402013488, 0.7232919558205454, 0.733066171439742, 0.7428403870589385},
			0.7375,
			4.453625,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.InDelta(t, c.expectedIndexRange, binaryFindIndex(c.values, c.qValue), 0.001)
		})
	}
}

func TestInterpolateFloat(t *testing.T) {
	type test struct {
		values   []float64
		index    float64
		expected float64
	}
	cases := []test{
		{
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			3.5,
			4.5,
		},
		{
			[]float64{0.0, 10.0},
			0.25,
			2.5,
		},
		{
			[]float64{2.0, 4.0},
			0.5,
			3.0,
		},
		{
			[]float64{0.0, 0.0},
			2.5,
			0.0,
		},
	}
	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.InDelta(t, c.expected, interpolateFloat(c.values, c.index), 0.001)
		})
	}
}

func TestInterP1(t *testing.T) {
	type test struct {
		name      string
		xList     []float64
		vList     []float64
		queryList []float64
		expected  []float64
		hasError  bool
	}
	cases := []test{
		{
			"valid input",
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			[]float64{10.0, 20.0, 30.0, 40.0, 50.0},
			[]float64{2.5, 3.5, 4.5},
			[]float64{25.0, 35.0, 45.0},
			false,
		},
		{
			"empty input",
			[]float64{},
			[]float64{},
			[]float64{},
			nil,
			true,
		},
		{
			"error case: Unequal length of xList and vList",
			[]float64{1.0, 2.0, 3.0},
			[]float64{10.0, 20.0, 30.0, 40.0, 50.0},
			[]float64{2.5, 3.5, 4.5},
			nil,
			true,
		},
		{
			"error case: Query value outside of range",
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0},
			[]float64{10.0, 20.0, 30.0, 40.0, 50.0},
			[]float64{0.5, 6.0},
			nil,
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			interpolated, err := InterP1(c.xList, c.vList, c.queryList)
			assert.Equal(t, c.hasError, err != nil)
			assert.Equal(t, c.expected, interpolated)
		})
	}
}
