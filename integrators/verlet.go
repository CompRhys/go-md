// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package integrators

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm
// for molecular dynamics using channels to provide simple parallelisarion.
func VerletStep(R, V []r3.Vector, L, M, dt float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan core.ForceReturn, N)
	for i := 0; i < N; i++ { go core.InternalForce(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.Index
		Fi := info.F
		A[i] = Fi.Mul(1.0/M)
		nR[i] = core.PutInBox(VerletNextR(R[i], V[i], A[i], dt), L)
	}
	for i := 0; i < N; i++ { go core.InternalForce(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.Index
		nFi := info.F
		nAi := nFi.Mul(1.0/M)
		nV[i] = VerletNextV(V[i], A[i], nAi, dt)
	}
	return nR, nV
}

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func VerletNextR(r, v, a r3.Vector, h float64) (nr r3.Vector) {
	nr = (r.Add(v.Mul(h))).Add(a.Mul(0.5*h*h))
	return 
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func VerletNextV(v, a1, a2 r3.Vector, h float64) (nv r3.Vector) {
	nv = v.Add((a1.Add(a2)).Mul(0.5*h))
	return 
}
