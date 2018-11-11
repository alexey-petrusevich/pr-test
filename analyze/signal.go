package analyze

type Signal struct {
	metaData     *MetaData
	binaryPoints *BinaryPoints
}

// .bin file metadata
type MetaData struct {
	Signature            string
	ChannelNumber        int32
	ChannelSize          int32
	SpectrumLineNumber   int32
	CutoffFrequency      int32
	FrequencyDefinition  float32
	DataBlockReceiveTime float32
	TotalReceiveTime     int32
	DataBlockNumber      int32
	DataSize             int32
	ReceivedBlocksNumber int32
	MaxValue             float32
	MinValue             float32
}

// .bin file points
type BinaryPoints struct {
	points []float32 // TODO: 32?
}
