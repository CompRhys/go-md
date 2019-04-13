package analysis

<<<<<<< HEAD
// import (
// 	"github.com/golang/geo/r3"
// 	"github.com/comprhys/moldyn/core"
//     "math"
//     "fmt"
// )
// func PrepareHistogram(r_max, L float64, dr int) []int, int{
//     if r_max > L/2 {
//         r_max = L/2
//         fmt.Println("r_max requested larger than L/2")
//     }
//     bins = r_max / dr
// }
// func UpdateHistogram(R []r3.Vector, r_max, L, dr float64, H []int) {
//     for i = 0; i<N-1; i++ {
//         for j = i+1; j<N; j++ {
//             r := Displacement(R[i], R[j], L)
//             if r < r_max {
//                 bin = int(r/dr)
//             }
//             H[bin] += 2
//         }
//     }
// }
// func NormaliseHistogram(dr, rho float64, bins, N int, H []int) []float64 {
//     rdf:= make([]float64, bins)
    
//     for i=0 ; i<bins; i++ {
//         r=dr*(i+0.5);
//         vol_bin=((i+1)*(i+1)*(i+1)-i*i*i)*dr*dr*dr;
//         nid=(4./3.)*math.Pi*vol_bin*rho;
//         rdf[i] = H[i]/(N*nid)
// }
=======
import (
	"github.com/golang/geo/r3"
	"github.com/comprhys/moldyn/core"
	
	"fmt"

	"math"
	"math/rand"

	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/gonum/stat/distuv"
)


func PrepareHistogram(r_max, L float64, dr int) []int, int{
    if r_max > L/2 {
        r_max = L/2
        fmt.Println("r_max requested larger than L/2")
    }
    bins = r_max / dr
}

func UpdateHistogram(R []r3.Vector, r_max, L, dr float64, H []int) {
    for i = 0; i<N-1; i++ {
        for j = i+1; j<N; j++ {
            r := Displacement(R[i], R[j], L)
            if r < r_max {
                bin = int(r/dr)
            }
            H[bin] += 2
        }
    }
}

func NormaliseHistogram(dr, rho float64, bins, N int, H []int) (rdf, rad []float64) {
    rdf:= make([]float64, bins)
    rad:= make([]float64, bins)
    
    for i=0 ; i<bins; i++ {
        r=dr*(i+0.5);
        vol_bin=((i+1)*(i+1)*(i+1)-i*i*i)*dr*dr*dr;
        nid=(4./3.)*math.Pi*vol_bin*rho;
        rdf[i] = H[i]/(N*nid)
	rad[i] = r
        
	return
}

func PlotHistogram(rad, rdf []float64) {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	
	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "r"
	p.Y.Label.Text = "g(r)"
	
	pts := make(plotter.XYs, len(rdf))
	
	pts.X = rad
	pts.Y = rad

	err = plotutil.AddLinePoints(p,
		"rdf", pts,
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
	
	
	

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
>>>>>>> ee75fa9ab66f2b34db0ee06e62d668fc848e3a68
