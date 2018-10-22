package core

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/integrators"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm
// for molecular dynamics using channels to provide simple parallelisarion.
func TimeStep(R, V []r3.Vector, L, M, h float64) ([]r3.Vector, []r3.Vector) {
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
