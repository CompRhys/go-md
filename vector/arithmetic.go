package vector

import "math"

// Sum calculates the vector sum of two vectors.
func Sum(a, b [3]float64) [3]float64 {
	r := [3]float64{0, 0, 0}
	for i, _ := range a {
		r[i] = a[i] + b[i]
	}
	return r
}

// Difference calculates the vector difference of two vectors.
func Difference(b, a [3]float64) [3]float64 {
	r := [3]float64{0, 0, 0}
	for i, _ := range a {
		r[i] = b[i] - a[i]
	}
	return r
}

// Scale calculates a vector scaled by a scalar.
func Scale(r [3]float64, s float64) [3]float64 {
	for i, v := range r {
		r[i] = v * s
	}
	return r
}

// Length calculates the length of a vector.
func Length(r [3]float64) float64 {
	var sum float64 = 0
	for _, v := range r {
		sum += v * v
	}
	return math.Sqrt(sum)
}

func SqLength(r [3]float64) float64 {
	var sum float64 = 0
	for _, v := range r {
		sum += v * v
	}
	return sum
}
