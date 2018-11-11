package cmd

import (
	"../analyze"
	"bufio"
	"io"
	"log"
	"os"
)

const (
	_DATA_OFFSET = 4 * 13
)

// METADATA constants
const (
	_SIGNATURE_SIZE = 4
	_INT_32_SIZE    = 4
	_FLOAT_32_SIZE  = 4
)

type CMD_RO struct {
	Executable
	*Command
}

func (command *CMD_RO) Execute(prevResult interface{}) (result interface{}) {
	if len(command.Command.Params) > 0 {
		path := command.Command.Params[0]
		var signals []analyze.Signal
		signals = append(signals, *getSignal(path))
		// return slice with signals
		return signals
	} else {
		return nil
	}
}

func (command *CMD_RO) Init(cmd string, position int, params []string) {
	command.Command = new(Command)
	command.Command.Init(cmd, position, params)
}

func getSignal(path string) *analyze.Signal {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sig := new(analyze.Signal)
	sig.MetaData = readMetaData(file)
	points := readDataPoints(file)
	for _, point := range points {
		sig.Points = append(sig.Points, point)
	}
	return sig
}

func readDataPoints(file *os.File) (points []float32) {
	reader := bufio.NewReader(file)
	// Seek to offset
	readN(reader, _DATA_OFFSET)
	//isReaded := false
	isReaded := false
	for !isReaded {
		bytes := readN(reader, _FLOAT_32_SIZE)
		if bytes != nil {
			point := float32(read_int32(bytes))
			points = append(points, point)
		} else {
			isReaded = true
		}
	}
	return points
}

func readMetaData(file *os.File) (metaData *analyze.MetaData) {
	metaData = new(analyze.MetaData)
	reader := bufio.NewReader(file)
	metaData.Signature = string(readN(reader, _SIGNATURE_SIZE))
	metaData.ChannelNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.ChannelSize = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.SpectrumLineNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.CutoffFrequency = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.FrequencyDefinition = float32(read_int32(readN(reader, _FLOAT_32_SIZE)))
	metaData.DataBlockReceiveTime = float32(read_int32(readN(reader, _FLOAT_32_SIZE)))
	metaData.TotalReceiveTime = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.DataBlockNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.DataSize = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.ReceivedBlocksNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.MaxValue = float32(read_int32(readN(reader, _FLOAT_32_SIZE)))
	metaData.MinValue = float32(read_int32(readN(reader, _FLOAT_32_SIZE)))
	return metaData
}

func read_int32(data []byte) int32 {
	return int32(uint32(data[0]) + uint32(data[1])<<8 + uint32(data[2])<<16 + uint32(data[3])<<24)
}

func readN(reader *bufio.Reader, n int) (result []byte) {
	for i := 0; i < n; i++ {
		byte, err := reader.ReadByte()
		if err == io.EOF {
			return result
		}
		result = append(result, byte)
	}
	return result
}
