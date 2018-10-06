// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package integrators

import (
	"github.com/golang/geo/r3"
)

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func NextR(r, v, a r3.Vector, h float64) (nr r3.Vector) {
	nr = (r.Add(v.Mul(h))).Add(a.Mul(0.5*h*h))
	// if brownian {
	// 	nr += 
	// }
	return 
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func NextV(v, a1, a2 r3.Vector, h float64) (nv r3.Vector) {
	nv = v.Add((a1.Add(a2)).Mul(0.5*h))
	// if brownian {
	// 	nv += 
	// }
	return 
}
