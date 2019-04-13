package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PlotTemperature(Temps []float64, dt float64) {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Radial Distribution Function"
	p.X.Label.Text = "t"
	p.Y.Label.Text = "T"

	bins := len(Temps)
	pts := make(plotter.XYs, bins)
	
	for i:=0 ; i<bins; i++ {
		pts[i].X = dt * float64(i)
		pts[i].Y = Temps[i]
	}

	err = plotutil.AddLinePoints(p,
		"rdf", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "temps.png"); err != nil {
		panic(err)
	}
}

func PlotHistogram(rad, rdf []float64) {
	// Make a plot and set its title.
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	bins := len(rad)

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
	