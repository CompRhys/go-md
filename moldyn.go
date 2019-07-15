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
	"github.com/comprhys/moldyn/integrators"
)

// Globals holds global simulation constants
type Globals struct {
	N           		int
	rho, M, T0, gamma, dt float64
	verlet bool
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	g := Globals{
		N: 8*8*8, rho: 0.8,
		T0: 1., M: 1.0,
		gamma: 1., dt: 0.01,
		verlet: true,
	}


	// current need N = M**3 

	V := float64(g.N) / g.rho
	L := math.Cbrt(V)

	Rs := core.InitPositionCubic(g.N, L)
	Vs := core.InitVelocity(g.N, g.T0, g.M)

	dr := L/120
	H, bins := analysis.PrepareHistogram(L/2, L, dr)
	r_max := dr * float64(bins)
	
	thermostat := integrators.PrepareLangevin(g.gamma, g.M, g.dt, g.T0, g.verlet)

	// Warm up
	for t := 0; t <= 100; t++ {
		// Rs, Vs = verlet.TimeStep(Rs, Vs, L, g.M, g.dt)
		Rs, Vs = integrators.LangevinStep(Rs, Vs, L, g.M, g.dt, thermostat)
	}

	T := 1000
	sample := 20
	var temps []float64
	start := time.Now()
	for t := 0; t <= T; t++ {
		// Rs, Vs = verlet.TimeStep(Rs, Vs, L, g.M, g.dt)
		Rs, Vs = integrators.LangevinStep(Rs, Vs, L, g.M, g.dt, thermostat)
		if t % sample == 0 {
			analysis.UpdateHistogram(Rs, r_max, L, dr, H)
			temps = append(temps, analysis.Temperature(Vs, g.M, g.N))
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("%v for %d time steps\n", elapsed, T)


	rdf, rad := analysis.NormaliseHistogram(dr, g.rho, bins, T, sample, g.N, H)
	plot.PlotHistogram(rad, rdf)
	plot.PlotTemperature(temps, g.dt)
}

