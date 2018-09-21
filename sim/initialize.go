// Package sim implements molecular dynamics simulation. The current implementation uses a Lennard Jones potential, but is generalizable to other potentials.
package sim

import (
	"github.com/golang/geo/r3"
	"math"
	"math/rand"
)

// InitPositionCubic initializes particle positions in a simple cubic configuration.
func InitPositionCubic(N int, L float64) []r3.Vector {
	R := make([]r3.Vector, N)
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
				pos := (r3.Vector{x, y, z}).Mul(rs)
				offset := r3.Vector{roffset, roffset, roffset}
				R[i] = pos.Sub(offset)
				i++
			}
		}
	}
	return R
}

// InitPositionFCC initializes particle positions in a face-centered cubic configuration
func InitPositionFCC(N int, L float64) []r3.Vector {
	R := make([]r3.Vector, N)
	Ncube := 1
	for N > 4*Ncube*Ncube*Ncube {
		Ncube++
	}
	o := -L / 2
	origin := r3.Vector{o, o, o}
	rs := L / float64(Ncube)
	roffset := rs / 2
	i := 0
	for x := 0; x < Ncube; x++ {
		x := float64(x)
		for y := 0; y < Ncube; y++ {
			y := float64(y)
			for z := 0; z < Ncube; z++ {
				z := float64(z)
				pos := (r3.Vector{x, y, z}).Mul(rs)
				pos = pos.Add(origin)
				R[i] = pos
				i++
				R[i] = pos.Add(r3.Vector{roffset, roffset, 0})
				i++
				R[i] = pos.Add(r3.Vector{roffset, 0, roffset})
				i++
				R[i] = pos.Add(r3.Vector{0, roffset, roffset})
				i++
			}
		}
	}
	return R
}

// InitVelocity initializes particle velocities selected from a random distribution.
// Ensures that the net momentum of the system is zero and scales the average kinetic energy to match a given temperature.
func InitVelocity(N int, T0 float64, M float64) []r3.Vector {
	V := make([]r3.Vector, N)
	rand.Seed(1)
	netP := r3.Vector{0, 0, 0}
	netE := 0.0
	for n := 0; n < N; n++ {
		newP := (r3.Vector{rand.Float64() , rand.Float64(), rand.Float64()}).Sub(r3.Vector{0.5,0.5,0.5})
		netP = netP.Add(newP)
		netE += newP.Norm2()
		V[n] = newP
	}
	netP = netP.Mul(1.0/float64(N))
	vscale := math.Sqrt(3.0 * float64(N) * T0 / (M * netE))
	for i, v := range V {
		correctedV := (v.Sub(netP)).Mul(vscale)
		V[i] = correctedV
	}
	return V
}
