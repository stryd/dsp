package dsp

import (
	"math"
	"reflect"
	"testing"
)

func TestLeastSquares(t *testing.T) {
	type testCase struct {
		points        []Point
		wantSlope     float64
		wantIntercept float64
	}
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
				{1, 1},
				{2, 2},
				{3, 5},
				{4, 10},
			}, 3, -3,
		},
		testCase{
			[]Point{
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
			}, 1, 0,
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

func TestPolynomial(t *testing.T) {
	type testCase struct {
		x      []float64
		y      []float64
		degree int
		want   []float64
	}
	cases := []testCase{
		testCase{
			[]float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321},
			2,
			[]float64{1, 2, 3},
		},
	}
	for _, c := range cases {
		dense, err := Polynomial(c.x, c.y, c.degree)
		if err != nil {
			t.Errorf("Failed the polynomial regression test: %v", err)
		}
		coeffs := make([]float64, c.degree+1)
		for i := 0; i < c.degree+1; i++ {
			coeffs[i] = math.Round(dense.At(i, 0))
		}
		if !reflect.DeepEqual(c.want, coeffs) {
			t.Errorf("Polynomial regression result is wrong, want %v, got %v", c.want, coeffs)
		}
	}
}

/*
func TestPolynomialRealData(t *testing.T) {
	file, err := os.Open("./poly_data")
	if err != nil {
		t.Errorf("Failed to read data file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	reListStr := lines[0]
	gradeListStr := lines[1]
	reStr := strings.Split(reListStr, ",")
	gradeStr := strings.Split(gradeListStr, ",")
	res := make([]float64, 0)
	grades := make([]float64, 0)
	for i := 0; i < len(reStr); i++ {
		if re, err := strconv.ParseFloat(reStr[i], 64); err == nil {
			res = append(res, re)
		}

		if grade, err := strconv.ParseFloat(gradeStr[i], 64); err == nil {
			grades = append(grades, grade)
		}
	}

	res = res[2000:2100]
	grades = grades[2000:2100]
	fmt.Println(res, grades)
	degree := 2
	dense, err := Polynomial(grades, res, degree)
	if err != nil {
		t.Errorf("Failed the polynomial regression test: %v", err)
		return
	}
	coeffs := make([]float64, degree+1)
	for i := 0; i < degree+1; i++ {
		coeffs[i] = math.Round(dense.At(i, 0))
	}
	fmt.Println(coeffs)
}
*/
