package fitprotocol

import (
	"bufio"
	"io"
)

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
	Reader io.Reader
	Buffer *bufio.Reader
}

func NewFitReader(reader io.Reader) (FitReader, error) {
	return FitReader{Reader: reader, Buffer: bufio.NewReader(reader)}, nil
}

func (s *FitReader) HeaderSize() (int, error) {
	// Should we just always grab the first byte of the file?
	// And basically ignore the buffer or reader current location?
	first_byte, peek_error := s.Buffer.Peek(1)
	if peek_error != nil {
		return 0, peek_error
	}
	return int(first_byte[0]), nil
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
