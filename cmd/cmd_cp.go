package cmd

import (
	"../analyze"
)

const _SIGNALS_PARAMS = "sig_params"

type CMD_CP struct {
	Executable
	*Command
}

func (command *CMD_CP) Execute(prevResult interface{}) (result interface{}) {
	commandResult := prevResult.(map[string]interface{})
	// TODO analyze other signals
	signals := commandResult[_RESULT_SIG].([]analyze.Signal)
	var params []analyze.SignalParameters
	for _, signal := range signals {
		sigParams := analyze.CalculateSignalParameters(&signal, int(signal.MetaData.ChannelSize))
		params = append(params, *sigParams)
	}
	commandResult[_SIGNALS_PARAMS] = params
	println("Min")
	println(params[0].MinValue)
	println("Max")
	println(params[0].MaxValue)
	println("SD")
	println(params[0].StandardDeviation)
	println("Peak")
	println(params[0].PeakFactor)
	println("Constant")
	println(params[0].Constant)
	println("Scope")
	println(params[0].Scope)
	return prevResult
}

func (command *CMD_CP) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}
