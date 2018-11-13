package cmd

import (
	"../analyze"
	"bufio"
	"log"
	"os"
)

type CMD_RG struct {
	Executable
	*Command
}

func (command *CMD_RG) Execute(prevResult interface{}) (result interface{}) {
	if len(command.Command.Params) > 0 {
		path := command.Command.Params[0]
		signalPaths, err := readLines(path)
		if err != nil {
			log.Fatal(err)
		}
		var signals []analyze.Signal
		for _, sigPath := range signalPaths {
			signal := *getSignal(sigPath)
			signals = append(signals, signal)
		}
		commandResult := make(map[string]interface{})
		commandResult[_RESULT_SIG] = signals
		// return slice with signals
		return commandResult
	} else {
		return nil
	}
}

func (command *CMD_RG) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
