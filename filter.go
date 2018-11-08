package dsp

// FilterByEWMA filters the signal using low pass filter based on exponential moving weighted moving average
func FilterByEWMA(signal []float64, alpha float64) []float64 {
	var newValues []float64
	filt := signal[0]
	for index := range signal {
		filt = alpha*filt + (1-alpha)*signal[index]
		newValues = append(newValues, filt)
	}

	return newValues
}

/***********************************************
 * IIR 2nd order filter used for trend analysis
 ***********************************************/

// Flips a vector of float
func flip(y []float64) []float64 {
	size := len(y)
	end := size - 1
	yy := make([]float64, len(y))
	copy(yy, y)
	var t float64
	for c := 0; c < size/2; c++ {
		t = yy[c]
		yy[c] = yy[end]
		yy[end] = t
		end--
	}
	return yy
}

// IIR 2nd order Filter
func filter(b, a, y []float64) []float64 {
	size := len(y)
	yy := make([]float64, size)
	df := []float64{0, 0}
	var xn, yn, s1, s2 float64
	for i := 0; i < size; i++ {
		xn = y[i]
		yn = b[0]*xn + df[0]
		s1 = df[1] + b[1]*xn - a[1]*yn
		s2 = b[2]*xn - a[2]*yn
		df[0] = s1
		df[1] = s2
		yy[i] = yn
	}
	return yy
}

// Filtfilt filters the signal using IIR 2nd order zero phase filter
func Filtfilt(b, a, y []float64) []float64 {
	var yy []float64
	yy = append(yy, y...)
	yy = append(yy, flip(y)...)

	// Filter forwards
	yy = filter(b, a, yy)
	// Reverse time
	yy = flip(yy)
	// Filter backwards
	yy = filter(b, a, yy)
	// Re-reverse time back to normal
	yy = flip(yy)

	return yy[:len(y)]
}
