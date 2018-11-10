package cmd

import (
	"../sync"
	"bufio"
	"fmt"
	"os"
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
		fmt.Printf("Your input: %s\n", input)
	}
	syncEntity.DoneExecuting()
}
