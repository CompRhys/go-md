package analysis

import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
	
	"fmt"

	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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
        rdf[i] = H[i]/(N_f*nid)
		rad[i] = r
		j += 1.
	}
	return
}

func PlotHistogram(rad, rdf []float64, bins int) {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	
	p.Title.Text = "Radial Distribution Function"
	p.X.Label.Text = "r"
	p.Y.Label.Text = "g(r)"
	
	pts := make(plotter.XYs, len(rdf))
	
	for i:=0 ; i<bins; i++ {
		pts[i].X = rad[i]
		pts[i].Y = rdf[i]
	}

	err = plotutil.AddLinePoints(p,
		"rdf", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
	