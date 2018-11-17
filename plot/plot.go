package plot

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func SaveDataToHtml(data map[float64]float64, xName, yName, signalName string) {
	var dataString string
	var keys []float64
	for k := range data {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	// To perform the opertion you want
	for _, key := range keys {
		dataString += getPointLikeArray(key, data[key])
	}

	//for key, value := range data {
	//	dataString += getPointLikeArray(key, value)
	//}
	dataString = "[" + dataString + "]"
	page := getFileDataAsString("./plot/html-output.template")
	output := strings.Replace(page, "%%data", dataString, -1)
	file, err := os.Create(signalName + ".html")
	if err != nil {
		panic(err)
	}
	file.WriteString(output)
}

func getFileDataAsString(path string) (result string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func getPointLikeArray(x float64, y float64) string {
	return "[" + fmt.Sprintf("%f", x) + "," + fmt.Sprintf("%f", y) + "],"
}

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
