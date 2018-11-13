package analyze

import "math"

type SignalParameters struct {
	MaxValue          float32
	MinValue          float32
	Scope             float32
	Constant          float32
	StandardDeviation float32
	PeakFactor        float32
}

func CalculateSignalParameters(signal *Signal) (parameters *SignalParameters) {
	parameters = new(SignalParameters)
	points := signal.Points
	parameters.MinValue = points[0]
	parameters.MaxValue = points[0]
	for i := 0; i < len(signal.Points); i++ {
		if parameters.MinValue > points[i] {
			parameters.MinValue = points[i]
		}
		if parameters.MaxValue < points[i] {
			parameters.MaxValue = points[i]
		}
	}
	parameters.Scope = parameters.MaxValue - parameters.MinValue
	for _, point := range points {
		parameters.Constant += point
	}
	parameters.Constant = parameters.Constant / float32(len(points))

	for _, point := range points {
		parameters.StandardDeviation += point * point
	}
	a := float64(parameters.StandardDeviation / float32(len(points)))
	parameters.StandardDeviation = float32(math.Sqrt(a))

	parameters.PeakFactor = max(parameters.MaxValue, parameters.MinValue) / parameters.StandardDeviation

	return parameters
}

func max(a, b float32) float32 {
	if a >= b {
		return a
	} else {
		return b
	}
}
