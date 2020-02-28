package dsp

import "gonum.org/v1/gonum/mat"

// Point holds the coordinates of the points used for regression
type Point struct {
	X float64
	Y float64
}

// LeastSquares returns implements the linear regression based on the method of least squares
// y = ax + b
func LeastSquares(points *[]Point) (slope float64, intercept float64) {
	n := float64(len(*points))

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumXX := 0.0

	for _, p := range *points {
		sumX += p.X
		sumY += p.Y
		sumXY += p.X * p.Y
		sumXX += p.X * p.X
	}

	base := (n*sumXX - sumX*sumX)
	slope = (n*sumXY - sumX*sumY) / base
	intercept = (sumXX*sumY - sumXY*sumX) / base

	return slope, intercept
}

func Polynomial(x, y []float64, degree int) (*mat.Dense, error) {
	a := vandermonde(x, degree)
	b := mat.NewDense(len(y), 1, y)
	c := mat.NewDense(degree+1, 1, nil)

	qr := new(mat.QR)
	qr.Factorize(a)

	err := qr.SolveTo(c, false, b)
	return c, err
}

func vandermonde(a []float64, degree int) *mat.Dense {
	x := mat.NewDense(len(a), degree+1, nil)
	for i := range a {
		for j, p := 0, 1.; j <= degree; j, p = j+1, p*a[i] {
			x.Set(i, j, p)
		}
	}
	return x
}
