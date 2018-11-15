package cmd

import (
	"../analyze"
)

const _SPECTRUM_POSTFIX = "-spectrum"

type CMD_SPECTRUM struct {
	Executable
	*Command
}

func (command *CMD_SPECTRUM) Execute(prevResult interface{}) (result interface{}) {
	commandResult := prevResult.(map[string]interface{})
	// TODO analyze other signals
	signals := commandResult[_RESULT_SIG].([]analyze.Signal)
	for _, signal := range signals {
		spectrum := analyze.CalculateSignalSpectrum(&signal)
		commandResult[signal.Name+_SPECTRUM_POSTFIX] = spectrum
	}
	return prevResult
}

func (command *CMD_SPECTRUM) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}
