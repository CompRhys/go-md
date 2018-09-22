package core

import (
	"math"
	"testing"
	"github.com/golang/geo/r3"
)

// Implicitly tests PutInBox also
func TestDisplacement(t *testing.T) {
	cases := []struct {
		a, b r3.Vector
		L    float64
		want r3.Vector
	}{
		{r3.Vector{0, 0, 0}, r3.Vector{1.0, 0, 0}, 2.0, r3.Vector{1.0, 0, 0}},
		{r3.Vector{-0.25, 0, 0}, r3.Vector{0.5, 0, 0}, 1.0, r3.Vector{-0.25, 0, 0}},
		{r3.Vector{0, 0, 0}, r3.Vector{1.0, 1.0, 1.0}, 1.0, r3.Vector{0, 0, 0}},
	}
	for _, c := range cases {
		got := Displacement(c.a, c.b, c.L)
		if got != c.want {
			t.Errorf("Displacement(%v, %v, %f) = %v, want %v", c.a, c.b, c.L, got, c.want)
		}
	}
}

// Implicitly tests Displacement and VectorLength also
func TestDistance(t *testing.T) {
	cases := []struct {
		a, b    r3.Vector
		L, want float64
	}{
		{r3.Vector{0, 0, 0}, r3.Vector{1.0, 0, 0}, 2.0, 1.0},
		{r3.Vector{-0.25, 0, 0}, r3.Vector{0.5, 0, 0}, 1.0, 0.25},
		{r3.Vector{0, 0, 0}, r3.Vector{1.0, 1.0, 1.0}, 1.0, 0},
		{r3.Vector{0, 0, 0}, r3.Vector{1.0, 1.0, 1.0}, 2.0, math.Sqrt(3)},
	}
	for _, c := range cases {
		got := Distance(c.a, c.b, c.L)
		if got != c.want {
			t.Errorf("Distance(%v, %v, %f) = %f, want %v", c.a, c.b, c.L, got, c.want)
		}
	}
}

