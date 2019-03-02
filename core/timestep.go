package core

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/integrators"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm
// for molecular dynamics using channels to provide simple parallelisarion.
func TimeStepVV(R, V []r3.Vector, L, M, h float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan ForceReturn, N)
	for i := 0; i < N; i++ { go InternalForce(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		Fi := info.F
		A[i] = Fi.Mul(1.0/M)
		nR[i] = PutInBox(integrators.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ { go InternalForce(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		nFi := info.F
		nAi := nFi.Mul(1.0/M)
		nV[i] = integrators.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}

// WarmUpStep as above but with force capping to prevent the system exploding
func WarmUpStepVV(R, V []r3.Vector, L, M, F_max, h float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan ForceReturn, N)
	for i := 0; i < N; i++ { go InternalForce(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		Fi := info.F
		F_mag := Fi.norm 
		if F_mag > F_max {
			Fi.Mul(F_max/F_mag)
		}
		A[i] = Fi.Mul(1.0/M)
		nR[i] = PutInBox(integrators.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ { go InternalForce(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		nFi := info.F
		nF_mag := nFi.norm 
		if nF_mag > F_max {
			nFi.Mul(F_max/nF_mag)
		}
		nAi := nFi.Mul(1.0/M)
		nV[i] = integrators.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}


// TimeStep evolves the system by one unit of time using the langevin thermostat
// for molecular dynamics using channels to provide simple parallelisarion.
func TimeStepBD(R, V []r3.Vector, L, M, h float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan ForceReturn, N)
	for i := 0; i < N; i++ { go InternalForce(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		Fi := info.F
		A[i] = Fi.Mul(1.0/M)
		nR[i] = PutInBox(integrators.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ { go InternalForce(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		nFi := info.F
		nAi := nFi.Mul(1.0/M)
		nV[i] = integrators.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}
