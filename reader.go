package fitprotocol

import (
	"bufio"
	"os"
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
	File   *os.File
	Buffer *bufio.Reader
}

func NewFitReader(file *os.File) (FitReader, error) {
	return FitReader{File: file, Buffer: bufio.NewReader(file)}, nil
}

func (s *FitReader) HeaderSize() (int, error) {
	header_size := make([]byte, 1)
	_, read_err := s.File.ReadAt(header_size, 0)
	if read_err != nil {
		return 0, read_err
	}
	return int(header_size[0]), nil
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
