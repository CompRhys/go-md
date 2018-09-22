// Package space implements utilities to do vector math with periodic boundary conditions.
package space

import "github.com/golang/geo/r3"

// PutInBox normalizes a vector to fit inside a cell with periodic boundary conditions.
func PutInBox(r r3.Vector, L float64) r3.Vector {
    switch {
        case r.X < -L/2:
            r.X = r.X + L
        case r.X > L/2:
            r.X = r.X - L
    }
    switch {
        case r.Y < -L/2:
            r.Y = r.Y + L
        case r.Y > L/2:
            r.Y = r.Y - L
    }
    switch {
        case r.Z < -L/2:
            r.Z = r.Z + L
        case r.Z > L/2:
            r.Z = r.Z - L
    }   
    return r
}

// Displacement calculates the smallest vector pointing from a to b in a cell with periodic boundary conditions.
func Displacement(a, b r3.Vector, L float64) r3.Vector {
    r := b.Sub(a)
    return PutInBox(r, L)
}

// Distance calculates the scalar distance between two points in a cell with periodic boundary conditions.
func Distance(a, b r3.Vector, L float64) float64 {
    d := Displacement(a, b, L)
    return d.Norm()
}

