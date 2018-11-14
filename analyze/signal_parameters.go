package analyze

import (
	"golang.org/x/exp/errors/fmt"
	"math"
	"strings"
)

type SignalParameters struct {
	MaxValue          float64
	MinValue          float64
	Scope             float64
	Constant          float64
	StandardDeviation float64
	PeakFactor        float64
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

	parameters.Constant /= float64(selectionSize)

	for i := 0; i < selectionSize && i < len(signal.Points); i++ {
		parameters.StandardDeviation += points[i] * points[i]
	}

	a := float64(parameters.StandardDeviation / float64(selectionSize))
	parameters.StandardDeviation = float64(math.Sqrt(a))

	parameters.PeakFactor = float64(math.Max(parameters.MaxValue, parameters.MinValue)) / parameters.StandardDeviation
	println(parameters)
	return parameters
}

func (params *SignalParameters) String() string {
	builder := new(strings.Builder)
	builder.WriteString("MaxValue: ")
	builder.WriteString(fmt.Sprint(params.MaxValue))
	builder.WriteString("\n")

	builder.WriteString("MinValue: ")
	builder.WriteString(fmt.Sprint(params.MinValue))
	builder.WriteString("\n")

	builder.WriteString("Scope: ")
	builder.WriteString(fmt.Sprint(params.Scope))
	builder.WriteString("\n")

	builder.WriteString("Constant: ")
	builder.WriteString(fmt.Sprint(params.Constant))
	builder.WriteString("\n")

	builder.WriteString("Standard deviation: ")
	builder.WriteString(fmt.Sprint(params.StandardDeviation))
	builder.WriteString("\n")

	builder.WriteString("Peak-factor: ")
	builder.WriteString(fmt.Sprint(params.PeakFactor))
	builder.WriteString("\n")
	return builder.String()
}
