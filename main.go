package main

import (
	"./cmd"
	"./sync"
	"fmt"
)

//func main() {
//	file, err := os.Open("./example_data/1.bin")
//	if err != nil {
//		log.Fatalf("First fatal: %s", err)
//	}
//	signal, err := wav.New(file)
//	if err != nil {
//		log.Fatalf("Second fatal: %s", err)
//	}
//	fmt.Println(signal.SampleRate)
//}

func main() {
	// example commands
	fmt.Println("ro=./example_data/1.bin exit")
	fmt.Println("rg=./example_data/text.txt exit")
	fmt.Println("ro=./example_data/1.bin spectrum cp save=wav,spectrum,html exit")
	syncEntity := sync.GetSyncEntity(1)
	cmd.WaitCommands(syncEntity)
	syncEntity.WaitAll()
}
