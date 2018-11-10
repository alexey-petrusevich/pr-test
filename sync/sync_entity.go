package sync

import (
	"log"
	"sync"
)

type SyncEntity struct {
	waitGroupSize int
	waitGroup     *sync.WaitGroup
}

func GetSyncEntity(waitGroupSize int) *SyncEntity {
	entity := new(SyncEntity)
	entity.waitGroupSize = waitGroupSize
	entity.waitGroup = new(sync.WaitGroup)
	entity.waitGroup.Add(waitGroupSize)
	return entity
}

func (entity *SyncEntity) DoneExecuting() {
	entity.waitGroup.Done()
}

func (entity *SyncEntity) WaitAll() {
	entity.waitGroup.Wait()
}

func (entity *SyncEntity) Error(err error) {
	log.Fatal(err)
	entity.waitGroup.Done()
}
