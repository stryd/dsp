package dsp

import (
	"fmt"
	"math"
	"sort"
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

/**
 * InterP1 simulate the interp1 method in matlab.
 * It returns interpolated values of a 1-D function at specific query points using linear interpolation.
 *
 * This is ported from the js code which is ported from the matlab code:
 * https://www.npmjs.com/package/interp1?activeTab=code
 *
 * @param  xList        Array of independent sample points.
 * @param  vList        Array of dependent values v(x) with length equal to xList.
 * @param  qeuryList    Array of coordinates of the query points.
 * @return              Interpolated values with length equal to queryList.
 */
func InterP1(xList, vList []float64, queryList []float64) ([]float64, error) {
	if len(xList) != len(vList) {
		return nil, fmt.Errorf("arrays of sample points xList and corresponding values vList have to have equal length.: %d vs. %d\n", len(xList), len(vList))
	}
	size := len(xList)
	if size == 0 {
		return nil, fmt.Errorf("arrays of sample points xList and corresponding values vList have to have length > 0.")
	}

	type Zip struct {
		X, V float64
	}
	// Combine x and v
	p := make([]Zip, size)
	for i := range xList {
		p[i] = Zip{xList[i], vList[i]}
	}

	// Sort asc
	sort.Slice(p, func(i, j int) bool {
		return p[i].X < p[j].X
	})

	// Check for double x values
	for i := range p {
		if i == 0 {
			continue
		}
		if p[i].X == p[i-1].X {
			return nil, fmt.Errorf("two sample points have equal value %f. This is not allowed.", p[i].X)
		}
	}

	// Split
	sortedX, sortedV := make([]float64, size), make([]float64, size)
	for _, v := range p {
		sortedX = append(sortedX, v.X)
		sortedV = append(sortedV, v.V)
	}

	// Interpolate
	r := []float64{}
	for _, xq := range queryList {
		// Determine index of range of query value.
		index := binaryFindIndex(sortedX, xq)

		// Check if value lies in interpolation range.
		if index == -1.0 {
			return nil, fmt.Errorf("query value %f lies outside of range. Extrapolation is not supported.", xq)
		}

		r = append(r, interpolateFloat(sortedV, index))
	}

	return r, nil
}

/**
 * interpolateFloat applies linear interpolation on a value based on index.
 * The index is allowed to be an non-integer. Used by InterP1D.
 *
 * This is ported from the js code which is ported from the matlab code:
 * https://www.npmjs.com/package/interp1?activeTab=code
 *
 * @param  values Array of values to interpolate between.
 * @param  index  Index of new to be interpolated value.
 * @return        Interpolated value.
 */
func interpolateFloat(values []float64, index float64) float64 {
	if index > float64(len(values)-1) {
		return 0
	}
	prev := math.Floor(index)
	next := math.Ceil(index)
	lambda := index - prev
	return (1-lambda)*values[int(prev)] + lambda*values[int(next)]
}

/**
 * Finds the index of range in which a query value is included in a sorted
 * array with binary search. Used by InterP1D.
 *
 * This is ported from the js code which is ported from the matlab code:
 * https://www.npmjs.com/package/interp1?activeTab=code
 *
 * @param   values Array sorted in ascending order.
 * @param   qValue Query value.
 * @return         Index of range plus percentage to next index.
 */
func binaryFindIndex(values []float64, qValue float64) float64 {
	// Special case of only one element in array.
	if len(values) == 1 && values[0] == qValue {
		return 0.0
	}

	// Determine bounds.
	lower := 0
	upper := len(values) - 1

	// Find index of range.
	for lower < upper {
		// Determine test range.
		mid := math.Floor(float64(lower+upper) / 2.0)
		prev := values[int(mid)]
		next := values[int(mid)+1]

		if qValue < prev {
			// Query value is below range.
			upper = int(mid)
		} else if qValue > next {
			// Query value is above range.
			lower = int(mid) + 1
		} else {
			// Query value is in range.
			return mid + (qValue-prev)/(next-prev)
		}
	}

	// Range not found.
	return -1.0
}
