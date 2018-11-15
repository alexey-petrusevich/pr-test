package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func SaveSpectrum(signalSpectrumKey string, spectrum map[int]float64, border int32) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = signalSpectrumKey
	p.X.Label.Text = "Freq"
	p.Y.Label.Text = "Amplit"

	pts := make(plotter.XYs, len(spectrum))
	for key, value := range spectrum {
		if key < int(border/2) {
			pts[key].X = float64(key)
			pts[key].Y = value
			if value > 50 {
				println(value)
			}
		}
	}

	err = plotutil.AddLinePoints(p,
		"Spectrum", pts)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(7*vg.Inch, 7*vg.Inch, signalSpectrumKey+".png"); err != nil {
		panic(err)
	}
}
