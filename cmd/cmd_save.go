package cmd

import (
	"../analyze"
	"../plot"
	"../wav"
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
	case "spectrum":
		saveSpectrum(commandResult)
	}
}

func saveDataToHtml(xys map[float64]float64, xName, yName, title string) {
	plot.SaveDataToHtml(xys, xName, yName, title)
}

func saveSpectrum(commandResult map[string]interface{}) {
	signals := commandResult[_RESULT_SIG].([]analyze.Signal)
	for _, signal := range signals {
		signalSpectrumKey := signal.Name + _SPECTRUM_POSTFIX
		spectrum := commandResult[signalSpectrumKey].(map[int]float64)
		plot.SaveSpectrum(signalSpectrumKey, spectrum, signal.MetaData.ChannelSize)

		newMap := translateMap(spectrum)
		saveDataToHtml(newMap, "Amp", "Freq", signal.Name)
	}
}

func saveWav(commandResult map[string]interface{}) {
	// TODO Save only signals?
	signals := commandResult[_RESULT_SIG].([]analyze.Signal)
	for _, signal := range signals {
		data := signal.Points
		wav.WriteWAVForSignal(signal.Name, 1, data)
	}
}

func translateMap(oldMap map[int]float64) (newMap map[float64]float64) {
	newMap = make(map[float64]float64)
	for key, value := range oldMap {
		newMap[float64(key)] = value
	}
	return newMap
}
