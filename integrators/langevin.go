// Package verlet implements velocity verlet for stepwise Newtonian mechanics.
package integrators

import (
	// "fmt"
	"math"
	// "os"
	"math/rand"
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
)

type VecReturn struct {
	Index int
	nV r3.Vector
}

type Thermostat struct {
	d float64
	q float64
	sigma float64
	verlet bool
}

func PrepareLangevin(gamma, M, h, T0 float64, verlet bool) (Thermostat) {
	d := h/2./M
	q := gamma*h/2
	sigma := math.Sqrt(h*T0*gamma/M)
	return Thermostat{d, q, sigma, verlet}
}


// Implementation of first order langevin dynamics
func LangevinStep(R, V []r3.Vector, L, M, dt float64, therm Thermostat) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	nR := make([]r3.Vector, N)
	Fi := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	cF := make(chan core.ForceReturn, N)

	for i := 0; i < N; i++ { go core.InternalForce(i, R, L, cF) }
	for i := 0; i < N; i++ {
		info := <- cF
		Fi[info.Index] = info.F
	}
	for n := 0; n < N; n++ { nV[n] = NextV(V[n], Fi[n], therm)}

	for n := 0; n < N; n++ {
		nR[n] = core.PutInBox(NextR(R[n], nV[n], dt), L)
	}

	for i := 0; i < N; i++ { go core.InternalForce(i, nR, L, cF) }
	for i := 0; i < N; i++ {
		info := <- cF
		Fi[info.Index] = info.F
	}
	for n := 0; n < N; n++ { nV[n] = NextV(nV[n], Fi[n], therm)}

	return nR, nV
}

// NextR calculates the next position vector based on current position, velocity, and acceleration.
func NextR(r, v r3.Vector, h float64) (nr r3.Vector) {
	nr = (r.Add(v.Mul(h)))
	return 
}

// NextV calculates the next velocity vector based on current velocity and acceleration and future acceleration.
func NextV(v, F r3.Vector, therm Thermostat) (nv r3.Vector) {
	
	d := therm.d
	q := therm.q
	sigma := therm.sigma

	nv = v.Add(F.Mul(d))

	if therm.verlet {
		eta := r3.Vector{rand.NormFloat64(), 
							rand.NormFloat64(), 
							rand.NormFloat64()}

		nv = nv.Sub(v.Mul(q)).Add(eta.Mul(sigma))	
	}

	return
}
