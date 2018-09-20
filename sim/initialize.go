// Package sim implements molecular dynamics simulation. The current implementation uses a Lennard Jones potential, but is generalizable to other potentials.
package sim

import (
	"github.com/quells/LennardJonesGo/vector"
	"math"
	"math/rand"
)

// InitPositionCubic initializes particle positions in a simple cubic configuration.
func InitPositionCubic(N int, L float64) [][3]float64 {
	R := make([][3]float64, N)
	Ncube := 1
	for N > Ncube*Ncube*Ncube {
		Ncube++
	}
	rs := L / float64(Ncube)
	roffset := (L - rs) / 2
	i := 0
	for x := 0; x < Ncube; x++ {
		x := float64(x)
		for y := 0; y < Ncube; y++ {
			y := float64(y)
			for z := 0; z < Ncube; z++ {
				z := float64(z)
				pos := vector.Scale([3]float64{x, y, z}, rs)
				offset := [3]float64{roffset, roffset, roffset}
				R[i] = vector.Difference(pos, offset)
				i++
			}
		}
	}
	return R
}

// InitPositionFCC initializes particle positions in a face-centered cubic configuration
func InitPositionFCC(N int, L float64) [][3]float64 {
	R := make([][3]float64, N)
	Ncube := 1
	for N > 4*Ncube*Ncube*Ncube {
		Ncube++
	}
	o := -L / 2
	origin := [3]float64{o, o, o}
	rs := L / float64(Ncube)
	roffset := rs / 2
	i := 0
	for x := 0; x < Ncube; x++ {
		x := float64(x)
		for y := 0; y < Ncube; y++ {
			y := float64(y)
			for z := 0; z < Ncube; z++ {
				z := float64(z)
				pos := vector.Scale([3]float64{x, y, z}, rs)
				pos = vector.Sum(pos, origin)
				R[i] = pos
				i++
				R[i] = vector.Sum(pos, [3]float64{roffset, roffset, 0})
				i++
				R[i] = vector.Sum(pos, [3]float64{roffset, 0, roffset})
				i++
				R[i] = vector.Sum(pos, [3]float64{0, roffset, roffset})
				i++
			}
		}
	}
	return R
}

// InitVelocity initializes particle velocities selected from a random distribution.
// Ensures that the net momentum of the system is zero and scales the average kinetic energy to match a given temperature.
func InitVelocity(N int, T0 float64, M float64) [][3]float64 {
	V := make([][3]float64, N)
	rand.Seed(1)
	netP := [3]float64{0, 0, 0}
	netE := 0.0
	for n := 0; n < N; n++ {
		for i := 0; i < 3; i++ {
			newP := rand.Float64() - 0.5
			netP[i] += newP
			netE += newP * newP
			V[n][i] = newP
		}
	}
	netP = vector.Scale(netP, 1.0/float64(N))
	vscale := math.Sqrt(3.0 * float64(N) * T0 / (M * netE))
	for i, v := range V {
		correctedV := vector.Scale(vector.Difference(v, netP), vscale)
		V[i] = correctedV
	}
	return V
}
