package analyze

import (
	"math"
)

type SignalParameters struct {
	MaxValue          float32
	MinValue          float32
	Scope             float32
	Constant          float32
	StandardDeviation float32
	PeakFactor        float32
}

func CalculateSignalParameters(signal *Signal, selectionSize int) (parameters *SignalParameters) {
	parameters = new(SignalParameters)
	points := signal.Points
	parameters.MinValue = points[0]
	parameters.MaxValue = points[0]
	for i := 0; i < selectionSize && i < len(signal.Points); i++ {
		if parameters.MinValue > points[i] {
			parameters.MinValue = points[i]
		}
		if parameters.MaxValue < points[i] {
			parameters.MaxValue = points[i]
		}
	}

	parameters.Scope = parameters.MaxValue - parameters.MinValue

	for i := 0; i < selectionSize && i < len(signal.Points); i++ {
		parameters.Constant += points[i]
	}

	parameters.Constant /= float32(selectionSize)

	for i := 0; i < selectionSize && i < len(signal.Points); i++ {
		parameters.StandardDeviation += points[i] * points[i]
	}

	a := float64(parameters.StandardDeviation / float32(selectionSize))
	parameters.StandardDeviation = float32(math.Sqrt(a))

	maxAbs := math.Abs(float64(parameters.MaxValue))
	minAbs := math.Abs(float64(parameters.MinValue))
	parameters.PeakFactor = float32(math.Max(maxAbs, minAbs)) / parameters.StandardDeviation

	return parameters
}

func max(a, b float32) float32 {
	if a >= b {
		return a
	} else {
		return b
	}
}
