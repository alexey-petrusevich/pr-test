package analyze

type Signal struct {
	Name     string
	MetaData *MetaData
	Points   []float64
}

// .bin file metadata
type MetaData struct {
	Signature            string
	ChannelNumber        int32
	ChannelSize          int32
	SpectrumLineNumber   int32
	CutoffFrequency      int32
	FrequencyDefinition  float64
	DataBlockReceiveTime float64
	TotalReceiveTime     int32
	DataBlockNumber      int32
	DataSize             int32
	ReceivedBlocksNumber int32
	MaxValue             float64
	MinValue             float64
}

// .bin file points
type BinaryPoints struct {
	// TODO: 32?
}
