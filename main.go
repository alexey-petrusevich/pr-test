package main

import (
	"./cmd"
	"./sync"
)

func main() {
	syncEntity := sync.GetSyncEntity(1)
	cmd.WaitCommands(syncEntity)
	syncEntity.WaitAll()
}
