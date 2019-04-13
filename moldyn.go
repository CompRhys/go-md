// LennardJones simulates molecular dynamics with the Lennard Jones potential.
package main

import (
	"fmt"
	"time"
	"runtime"
	"math"
	"math/rand"
	"github.com/comprhys/moldyn/core"
	"github.com/comprhys/moldyn/analysis"
	"github.com/comprhys/moldyn/plot"
	// "github.com/comprhys/moldyn/integrators/verlet"
	"github.com/comprhys/moldyn/integrators/langevin"
)

// Globals holds global simulation constants
type Globals struct {
	N           		int
	rho, M, T0, gamma, dt float64
}

func init() {
	rand.Seed(rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	g := Globals{
		N: 512, rho: 0.8,
		T0: 1., M: 1.0,
		gamma: .5, dt: 0.01,
	}

	V := float64(g.N) / g.rho
	L := math.Cbrt(V)

	Rs := core.InitPositionCubic(g.N, L)
	Vs := core.InitVelocity(g.N, g.T0, g.M)

	dr := L/120
	H, bins := analysis.PrepareHistogram(L/2, L, dr)
	r_max := dr * float64(bins) 

	T := 1000
	var temps []float64
	start := time.Now()
	for t := 0; t <= T; t++ {
		// Rs, Vs = verlet.TimeStep(Rs, Vs, L, g.M, g.dt)
		Rs, Vs = langevin.TimeStep(Rs, Vs, L, g.M, g.T0, g.gamma, g.dt)
		if t % 10 == 0 {
			analysis.UpdateHistogram(Rs, r_max, L, dr, H)
			temps = append(temps, analysis.Temperature(Vs, g.M, g.N))
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("%v for %d time steps\n", elapsed, T)

	rdf, rad := analysis.NormaliseHistogram(dr, g.rho, bins, g.N, H)
	plot.PlotHistogram(rad, rdf)
	plot.PlotTemperature(temps, g.dt)
}

