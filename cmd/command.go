package cmd

import (
	"fmt"
	"log"
	"strings"
)

const _INPUT_SPLITTER = " "
const _TOKEN_SPLITTER = "="
const _PARAMS_SPLITTER = ","

type Executable interface {
	Init(cmd string, position int, params []string)
	Execute(prevResult interface{}) (result interface{})
}

type Command struct {
	Command  string
	Position int
	Params   []string
}

func (command *Command) Execute(prevResult interface{}) (result interface{}) {
	// By default - nothing
	return nil
}

func (command *Command) Init(cmd string, position int, params []string) {
	command.Command = cmd
	command.Position = position
	command.Params = params
}

func CommandFactory(input string) []Executable {
	commandsTokens := tokenizeInput(input)
	if len(commandsTokens) == 0 {
		log.Println("Empty command line!")
		return nil
	}
	var commands []Executable

	for index, token := range commandsTokens {
		cmd, params := parseToken(token)
		command := getCommand(cmd)
		if command != nil {
			command.Init(token, index, params)
			commands = append(commands, command)
		} else {
			fmt.Printf("Invalid command - %s", cmd)
		}
	}
	return commands
}

func parseToken(token string) (cmd string, params []string) {
	pieces := strings.Split(token, _TOKEN_SPLITTER)
	cmd = pieces[0]
	if len(pieces) == 2 {
		paramsPieces := pieces[1]
		params = strings.Split(paramsPieces, _PARAMS_SPLITTER)
	}
	return cmd, params
}

func getCommand(commandToken string) Executable {
	var command Executable
	switch commandToken {
	case CMD_READ_GROUP:
		command = new(CMD_RG)
	case CMD_READ_ONE:
		command = new(CMD_RO)
	case CMD_EXIT_EXEC:
		command = new(CMD_EXIT)
	}
	return command
}

func tokenizeInput(input string) []string {
	return strings.Split(input, _INPUT_SPLITTER)
}
