package fitprotocol

import "bufio"

const (
	HEADER_CRC_SIZE        = 2
	DATA_CRC_SIZE          = 2
	COMPRESSED_HEADER_MASK = 0x80
	MSG_DEF_MASK           = 0x40
	MSG_HEADER_MASK        = 0x00
	LOCAL_MSG_NUM_MASK     = 0x0F
	DEV_DATA_MASK          = 0x20
)

type FitReader struct {
	Reader *bufio.Reader
}

func NewFitReader(reader *bufio.Reader) (FitReader, error) {
	return FitReader{Reader: reader}, nil
}

func (s *FitReader) HeaderSize() int {
	return 0
}

func (s *FitReader) DataSize() int {
	return 0
}

func (s *FitReader) HeaderCRCvalid() bool {
	return false
}

func (s *FitReader) DataCRCvalid() bool {
	return false
}

// Next Message func?

// Next Byte func
