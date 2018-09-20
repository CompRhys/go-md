package sim

import (
	"github.com/quells/LennardJonesGo/space"
	"github.com/quells/LennardJonesGo/vector"
	"github.com/quells/LennardJonesGo/verlet"
)

// TimeStep evolves the system by one unit of time using the Velocity Verlet algorithm for molecular dynamics.
func TimeStep(R, V [][3]float64, L, M, h float64) ([][3]float64, [][3]float64) {
	N := len(R)
	A := make([][3]float64, N)
	nR := make([][3]float64, N)
	nV := make([][3]float64, N)
	for i := 0; i < N; i++ {
		Fi := InternalForce(i, R, L)
		A[i] = vector.Scale(Fi, 1.0/M)
		nR[i] = space.PutInBox(verlet.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ {
		nFi := InternalForce(i, nR, L)
		nAi := vector.Scale(nFi, 1.0/M)
		nV[i] = verlet.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}

// TimeStepParallel does the same as TimeStep but with channels
func TimeStepParallel(R, V [][3]float64, L, M, h float64) ([][3]float64, [][3]float64) {
	N := len(R)
	A := make([][3]float64, N)
	nR := make([][3]float64, N)
	nV := make([][3]float64, N)
	c := make(chan ForceReturn, N)
	for i := 0; i < N; i++ { go InternalForceParallel(i, R, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		Fi := info.F
		A[i] = vector.Scale(Fi, 1.0/M)
		nR[i] = space.PutInBox(verlet.NextR(R[i], V[i], A[i], h), L)
	}
	for i := 0; i < N; i++ { go InternalForceParallel(i, nR, L, c) }
	for n := 0; n < N; n++ {
		info := <-c
		i := info.i
		nFi := info.F
		nAi := vector.Scale(nFi, 1.0/M)
		nV[i] = verlet.NextV(V[i], A[i], nAi, h)
	}
	return nR, nV
}
