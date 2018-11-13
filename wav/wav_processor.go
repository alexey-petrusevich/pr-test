package wav

import (
	"fmt"
	"github.com/cryptix/wav"
	"os"
)

const (
	DEFAULT_CHANNELS = 1     // 1
	DEFAULT_BITS     = 32    // Tone is lower when value is lower (8, 16, 32, 64 ...)
	DEFAULT_RATE     = 44100 // 44100 32768

	DEFAULT_OUT_PATH = "./"
)

func WriteWAVForSignal(name string, soundLength int, signal []float32) {
	meta := wav.File{
		Channels:        DEFAULT_CHANNELS,
		SampleRate:      DEFAULT_RATE,
		SignificantBits: DEFAULT_BITS,
	}
	WriteWAVByMeta(name, soundLength, signal, meta)
}

func WriteWAVByMeta(name string, soundLength int, signal []float32, meta wav.File) {
	wavOut, err := os.Create(DEFAULT_OUT_PATH + name + ".wav")
	checkErr(err)
	defer wavOut.Close()

	writer, err := meta.NewWriter(wavOut)
	checkErr(err)
	defer writer.Close()

	for n := 0; n < soundLength; n++ {
		for idx := range signal {
			funRes := int32(
				signal[idx],
			)
			writer.WriteInt32(funRes)
		}
		checkErr(err)
	}

	writer.Close()
	fmt.Println()
	fmt.Println("WAV file '" + name + "' created successful")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
