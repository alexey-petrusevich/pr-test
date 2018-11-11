package cmd

type CMD_RG struct {
	Executable
	*Command
}

func (command *CMD_RG) Execute(prevResult interface{}) (result interface{}) {
	return "RG"
}

func (command *CMD_RG) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}
