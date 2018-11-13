package cmd

import (
	"../analyze"
	"../wav"
	"strconv"
)

type CMD_SAVE struct {
	Executable
	*Command
}

func (command *CMD_SAVE) Execute(prevResult interface{}) (result interface{}) {
	commandResult := prevResult.(map[string]interface{})
	params := command.Params
	for _, param := range params {
		saveByParam(param, commandResult)
	}
	return commandResult
}

func (command *CMD_SAVE) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}

func saveByParam(saveType string, commandResult map[string]interface{}) {
	switch saveType {
	case "wav":
		saveWav(commandResult)
	}
}

func saveWav(commandResult map[string]interface{}) {
	// TODO Save only signals?
	signals := commandResult[_RESULT_SIG].([]analyze.Signal)
	for index, signal := range signals {
		data := signal.Points
		wav.WriteWAVForSignal(strconv.Itoa(index)+_RESULT_SIG, 1, data)
	}
}
