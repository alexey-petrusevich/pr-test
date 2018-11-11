package cmd

type CMD_RO struct {
	Executable
	*Command
}

func (command *CMD_RO) Execute(prevResult interface{}) (result interface{}) {

}

func (command *CMD_RO) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}
