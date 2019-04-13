package analysis

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
)

// PairwiseLennardJonesPotential calculates the Lennard Jones potential between two positions.
func PairwiseLennardJonesPotential(Ri, Rj r3.Vector, L float64) float64 {
	r := core.Displacement(Ri, Rj, L)
	R2 := r.Norm2()
	iR2 := 1.0 / R2
	iR6 := iR2 * iR2 * iR2
	iR12 := iR6 * iR6
	return 4 * (iR12 - iR6)
}

// KineticEnergy calculates the kinetic energy of a particle.
func KineticEnergy(v r3.Vector, m float64) float64 {
	v2 := v.Norm2()
	return 0.5 * m * v2
}

// TotalKineticEnergy calculates the kinetic energy of all particles in the system.
func TotalKineticEnergy(V []r3.Vector, m float64) (sum float64) {
	for _, v := range V {
		sum += KineticEnergy(v, m)
	}
	return
}

// Temperature calculates the temperature of the system from the average kinetic energy of the particles.
func Temperature(V []r3.Vector, m float64, N int) float64 {
	return TotalKineticEnergy(V, m) * 2 / 3 / float64(N)
}

// TotalPotentialEnergy calculates the potential energy of the system due to pairwise particle interactions.
func TotalPotentialEnergy(Rs []r3.Vector, L float64) (sum float64) {
	for i := 0; i < len(Rs)-1; i++ {
		for j := i + 1; j < len(Rs); j++ {
			sum += PairwiseLennardJonesPotential(Rs[i], Rs[j], L)
		}
	}
	return
}

// TotalEnergy calculates the total energy of the system.
func TotalEnergy(Rs, Vs []r3.Vector, L, M float64) (sum float64) {
	sum += TotalKineticEnergy(Vs, M)
	sum += TotalPotentialEnergy(Rs, L)
	return
}
