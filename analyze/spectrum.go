package analyze

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
)

func CalculateSignalSpectrum(signal *Signal) (result map[int]float64) {
	points := signal.Points[0:signal.MetaData.ChannelSize]
	fftSignalPoints := fft.IFFTReal(points)
	result = make(map[int]float64)
	for index, fftSig := range fftSignalPoints {
		result[index] = cmplx.Abs(fftSig) * 2
	}
	return result
}

func ditfft2(x []float64, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(x[0], 0)
		return
	}
	ditfft2(x, y, n/2, 2*s)
	ditfft2(x[s:], y[n/2:], n/2, 2*s)
	for k := 0; k < n/2; k++ {
		tf := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * y[k+n/2]
		y[k], y[k+n/2] = y[k]+tf, y[k]-tf
	}
}
