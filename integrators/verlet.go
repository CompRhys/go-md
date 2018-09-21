// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package verlet

import (
	"github.com/golang/geo/r3"
)

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func NextR(r, v, a r3.Vector, h float64) r3.Vector {
	return (r.Add(v.Mul(h))).Add(a.Mul(0.5*h*h))
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func NextV(v, a1, a2 r3.Vector, h float64) r3.Vector {
	return v.Add((a1.Add(a2)).Mul(0.5*h))
}
