package cmd

import (
	"os"
)

type CMD_EXIT struct {
	Executable
	*Command
}

func (command *CMD_EXIT) Execute(prevResult interface{}) (result interface{}) {
	os.Exit(100)
	return nil
}

func (command *CMD_EXIT) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}
