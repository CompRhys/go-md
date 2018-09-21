package sim

import (
	"fmt"

	"github.com/comprhys/moldyn/space"
	"github.com/golang/geo/r3"
)

// PairwiseLennardJonesForce calculates the force vector on particle Ri due to Rj using the Lennard Jones potential.
func PairwiseLennardJonesForce(Ri, Rj r3.Vector, L float64) r3.Vector {
	if Ri == Rj {
		panic(fmt.Sprintf("%v and %v are equal, the pairwise force is infinite", Ri, Rj))
	}
	r := space.Displacement(Ri, Rj, L)
	R2 := r.Norm2()
	iR2 := 1.0 / R2
	iR8 := iR2 * iR2 * iR2 * iR2
	iR14 := iR8 * iR2 * iR2 * iR2
	f := 4 * (-12*iR14 + 6*iR8)
	return r.Mul(f)
	// magR := vector.Length(r)
	// f := 4 * (-12*math.Pow(magR, -13) + 6*math.Pow(magR, -7))
	// return vector.Scale(r, f/magR)
}

// InternalForce calculates the total force vector on particle Ri due to the other particles in R due to a pairwise force.
func InternalForce(i int, R []r3.Vector, L float64) r3.Vector {
	F := r3.Vector{0, 0, 0}
	for j := range R {
		if i != j {
			F = F.Add(PairwiseLennardJonesForce(R[i], R[j], L))
		}
	}
	return F
}

// ForceReturn holds the index and force on a particle
type ForceReturn struct {
	i int
	F r3.Vector
}

// InternalForceParallel does the same as InternalForce but with channels
func InternalForceParallel(i int, R []r3.Vector, L float64, c chan ForceReturn) {
	F := r3.Vector{0, 0, 0}
	for j := range R {
		if i != j {
			F = F.Add(PairwiseLennardJonesForce(R[i], R[j], L))
		}
	}
	c <- ForceReturn{i, F}
}
