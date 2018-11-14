package analyze

import "github.com/mjibson/go-dsp/fft"

func CalculateSignalSpectrum(signal *Signal) {
	fftSignalPoints := fft.FFTReal(signal.Points)
	for _, fftSig := range fftSignalPoints {
		println(fftSig)
	}
}
