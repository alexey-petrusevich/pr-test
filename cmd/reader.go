package cmd

import (
	"../sync"
	"bufio"
	"os"
	"strings"
)

const _READER_DELIMETER = '\n'

// Wait commands
// Need run it in GoRoutine
func WaitCommands(syncEntity *sync.SyncEntity) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString(_READER_DELIMETER)
		if err != nil {
			syncEntity.Error(err)
		}
		commands := CommandFactory(strings.Trim(input, "\n"))
		var prevResult interface{}
		for _, cmd := range commands {
			prevResult = cmd.Execute(prevResult)
			//fmt.Printf("%i : %i\n", i, cmd.Execute(nil))
		}
	}
	syncEntity.DoneExecuting()
}
