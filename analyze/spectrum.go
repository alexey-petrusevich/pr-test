package analyze

import (
	"github.com/mjibson/go-dsp/fft"
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
