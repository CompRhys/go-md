// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package langevin

import (
	"math"
	"math/rand"
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm
// for molecular dynamics using channels to provide simple parallelisarion.
func TimeStep(R, V []r3.Vector, L, M, T0, gamma, dt float64) ([]r3.Vector, []r3.Vector) {
	hdt := dt/2
	N := len(R)
	hR := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan core.ForceReturn, N)
	for n := 0; n < N; n++ {
		hR[n] = core.PutInBox(NextR(R[n], V[n], hdt), L)
	}
	for i := 0; i < N; i++ { go core.InternalForce(i, hR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.Index
		Fi := info.F
		// Fi := core.InternalForce(n, hR, L)
		nV[i] = NextV(V[i], Fi, dt, gamma, T0, M)
	}
	for n := 0; n < N; n++ {
		nR[n] = core.PutInBox(NextR(hR[n], V[n], hdt), L)
	}

	return nR, nV
}

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func NextR(r, v r3.Vector, h float64) (nr r3.Vector) {
	nr = (r.Add(v.Mul(h)))
	return 
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func NextV(v, F r3.Vector, h, gamma, T0, M float64) (nv r3.Vector) {
	d := math.Exp(-gamma/M*h)
	q := M/gamma * (1 - d)
	sigma := math.Sqrt(M*T0*(1-d*d))
	eta := r3.Vector{rand.NormFloat64(),
				rand.NormFloat64(),
				rand.NormFloat64()}

	nv = v.Mul(d).Add(F.Mul(q)).Add(eta.Mul(sigma))
	return 
}
