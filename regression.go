package dsp

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
