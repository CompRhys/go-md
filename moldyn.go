// LennardJones simulates molecular dynamics with the Lennard Jones potential.
package main

import (
	"fmt"
	"time"
	"runtime"
	"github.com/comprhys/moldyn/core"
)

// Globals holds global simulation constants
type Globals struct {
	N           int
	L, M, T0, h float64
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	g := Globals{
		N: 2048, L: 12.6992084,
		T0: 0.2, M: 48.0,
		h: 0.01,
	}

	Rs := core.InitPositionFCC(g.N, g.L)
	Vs := core.InitVelocity(g.N, g.T0, g.M)

	T := 10
	start := time.Now()
	for t := 1; t <= T; t++ {
		Rs, Vs = core.TimeStep(Rs, Vs, g.L, g.M, g.h)
	}
	elapsed := time.Since(start)
	fmt.Printf("%v for %d time steps\n", elapsed, T)
}

