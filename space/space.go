// Package space implements utilities to do vector math with periodic boundary conditions.
package space

import "github.com/quells/LennardJonesGo/vector"

// PutInBox normalizes a vector to fit inside a cell with periodic boundary conditions.
func PutInBox(r [3]float64, L float64) [3]float64 {
	for i, v := range r {
		switch {
		case v < -L/2:
			r[i] = v + L
		case v > L/2:
			r[i] = v - L
		}
	}
	return r
}

// Displacement calculates the smallest vector pointing from a to b in a cell with periodic boundary conditions.
func Displacement(a, b [3]float64, L float64) [3]float64 {
	r := vector.Difference(b, a)
	return PutInBox(r, L)
}

// Distance calculates the scalar distance between two points in a cell with periodic boundary conditions.
func Distance(a, b [3]float64, L float64) float64 {
	d := Displacement(a, b, L)
	return vector.Length(d)
}

// PointsAreEqual tests whether two points are equal in a cell with periodic boundary conditions.
func PointsAreEqual(a, b [3]float64, L float64) bool {
	x := a[0] == b[0]
	y := a[1] == b[1]
	z := a[2] == b[2]
	return x && y && z
	//return Displacement(a, b, L) == [3]float64{0, 0, 0}
}
