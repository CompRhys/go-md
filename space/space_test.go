package space

import (
	"math"
	"testing"
)

// Implicitly tests PutInBox also
func TestDisplacement(t *testing.T) {
	cases := []struct {
		a, b [3]float64
		L    float64
		want [3]float64
	}{
		{[3]float64{0, 0, 0}, [3]float64{1.0, 0, 0}, 2.0, [3]float64{1.0, 0, 0}},
		{[3]float64{-0.25, 0, 0}, [3]float64{0.5, 0, 0}, 1.0, [3]float64{-0.25, 0, 0}},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 1.0, 1.0}, 1.0, [3]float64{0, 0, 0}},
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
		a, b    [3]float64
		L, want float64
	}{
		{[3]float64{0, 0, 0}, [3]float64{1.0, 0, 0}, 2.0, 1.0},
		{[3]float64{-0.25, 0, 0}, [3]float64{0.5, 0, 0}, 1.0, 0.25},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 1.0, 1.0}, 1.0, 0},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 1.0, 1.0}, 2.0, math.Sqrt(3)},
	}
	for _, c := range cases {
		got := Distance(c.a, c.b, c.L)
		if got != c.want {
			t.Errorf("Distance(%v, %v, %f) = %f, want %v", c.a, c.b, c.L, got, c.want)
		}
	}
}

func TestPointsAreEqual(t *testing.T) {
	cases := []struct {
		a, b [3]float64
		L    float64
		want bool
	}{
		{[3]float64{0, 0, 0}, [3]float64{1.0, 0, 0}, 2.0, false},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 0, 0}, 1.0, true},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 1.0, 1.0}, 1.0, true},
		{[3]float64{0, 0, 0}, [3]float64{1.0, 1.0, 1.0}, 2.0, false},
	}
	for _, c := range cases {
		got := PointsAreEqual(c.a, c.b, c.L)
		if got != c.want {
			t.Errorf("PointsAreEqual(%v, %v, %f) = %t, want %t", c.a, c.b, c.L, got, c.want)
		}
	}
}
