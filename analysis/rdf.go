package analysis

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
	
	"fmt"

	"math"
)


func PrepareHistogram(r_max, L, dr float64) (H []float64, bins int){
    if r_max > L/2 {
        r_max = L/2
        fmt.Println("r_max requested larger than L/2")
    }
	bins = int(r_max / dr)
	
	H = make([]float64, bins)

	return
}

func UpdateHistogram(R []r3.Vector, r_max, L, dr float64, H []float64) {
	N := len(R)
	
	for i := 0; i<N-1; i++ {
        for j := i+1; j<N; j++ {
            r := core.Distance(R[i], R[j], L)
            if r < r_max {
				bin := int(r/dr)
				H[bin] += 2.
            }
        }
    }
}

func NormaliseHistogram(dr, rho float64, bins, N int, H []float64) (rdf, rad []float64) {
    rdf = make([]float64, bins)
    rad = make([]float64, bins)
	
	N_f := float64(N)
	j := 0.
    for i:=0 ; i<bins; i++ {
		
        r:=dr*(j+0.5);
        vol_bin:=((j+1)*(j+1)*(j+1)-j*j*j)*dr*dr*dr;
        nid:=(4./3.)*math.Pi*vol_bin*rho;
        rdf[i] = H[i]/N_f/nid
		rad[i] = r
		j += 1.
	}
	return
}

