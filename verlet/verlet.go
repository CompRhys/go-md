// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package verlet

import (
	"github.com/quells/LennardJonesGo/vector"
)

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func NextR(r, v, a [3]float64, h float64) [3]float64 {
	return vector.Sum(vector.Sum(r, vector.Scale(v, h)), vector.Scale(a, 0.5*h*h))
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func NextV(v, a1, a2 [3]float64, h float64) [3]float64 {
	return vector.Sum(v, vector.Scale(vector.Sum(a1, a2), 0.5*h))
}
