package core

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/space"
	"github.com/comprhys/moldyn/integrators"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm for molecular dynamics.
func TimeStep(R, V []r3.Vector, L, M, h float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	for i := 0; i < N; i++ {
		Fi := InternalForce(i, R, L)
		A[i] = Fi.Mul(1.0/M)
		nR[i] = space.PutInBox(verlet.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ {
		nFi := InternalForce(i, nR, L)
		nAi := nFi.Mul(1.0/M)
		nV[i] = verlet.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}

// TimeStepParallel does the same as TimeStep but with channels
func TimeStepParallel(R, V []r3.Vector, L, M, h float64) ([]r3.Vector, []r3.Vector) {
	N := len(R)
	A := make([]r3.Vector, N)
	nR := make([]r3.Vector, N)
	nV := make([]r3.Vector, N)
	c := make(chan ForceReturn, N)
	for i := 0; i < N; i++ { go InternalForceParallel(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		Fi := info.F
		A[i] = Fi.Mul(1.0/M)
		nR[i] = space.PutInBox(verlet.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ { go InternalForceParallel(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		nFi := info.F
		nAi := nFi.Mul(1.0/M)
		nV[i] = verlet.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}
