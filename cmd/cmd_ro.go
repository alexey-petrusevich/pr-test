package cmd

import (
	"../analyze"
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"math"
	"os"
)

const (
	_DATA_OFFSET = 4 * 13
	// .bs here, it's flag that signal it's a base
	_BASE_SIGNAL_POSTFIX = ".bs"
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
		commandResult := make(map[string]interface{})
		commandResult[_RESULT_SIG] = signals
		// return slice with signals
		return commandResult
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
	sig.Name = createSignalName(file)
	sig.MetaData = readMetaData(file)
	points := readDataPoints(file)
	for _, point := range points {
		sig.Points = append(sig.Points, point)
	}
	return sig
}

func readDataPoints(file *os.File) (points []float64) {
	reader := bufio.NewReader(file)
	// Seek to offset
	readN(reader, _DATA_OFFSET)
	//isReaded := false
	isReaded := false
	for !isReaded {
		bytes := readN(reader, _FLOAT_32_SIZE)
		if bytes != nil {
			point := read_float32(bytes)
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
	metaData.FrequencyDefinition = read_float32(readN(reader, _FLOAT_32_SIZE))
	metaData.DataBlockReceiveTime = read_float32(readN(reader, _FLOAT_32_SIZE))
	metaData.TotalReceiveTime = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.DataBlockNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.DataSize = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.ReceivedBlocksNumber = int32(read_int32(readN(reader, _INT_32_SIZE)))
	metaData.MinValue = read_float32(readN(reader, _FLOAT_32_SIZE))
	metaData.MaxValue = read_float32(readN(reader, _FLOAT_32_SIZE))
	return metaData
}

func read_float32(data []byte) float64 {
	bits := binary.LittleEndian.Uint32(data)
	float := math.Float32frombits(bits)
	return float64(float)
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

func createSignalName(file *os.File) string {
	return file.Name() + _BASE_SIGNAL_POSTFIX
}
