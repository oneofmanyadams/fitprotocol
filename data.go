package fitprotocol

import (
	"bufio"
	"errors"
)

const (
	DATA_CRC_SIZE          = 2
	COMPRESSED_HEADER_MASK = 0x80
	MSG_DEF_MASK           = 0x40
	MSG_HEADER_MASK        = 0x00
	LOCAL_MSG_NUM_MASK     = 0x0F
	DEV_DATA_MASK          = 0x20
)

var BYTE_READ_ERROR = errors.New("Error attempting to read from byte buffer.")

type Data struct {
	Reader     *bufio.Reader
	HeaderSize int
	DataSize   int
	Position   int
	ReadLines  int
	ReadLimit  int
}

func NewData(reader *bufio.Reader, header_size, data_size int) Data {
	return Data{
		Reader:     reader,
		HeaderSize: header_size,
		DataSize:   data_size,
		ReadLimit:  data_size}
}

func (s *Data) Read() (byte, error) {
	b, err := s.Reader.ReadByte()
	if err != nil {
		return 0x0, nil
	}
	s.ReadLines++
	s.Position++
	return b, nil
}

func (s *Data) ReadRemaining() ([]byte, error) {
	var bytes []byte
	for s.ReadLines < s.DataSize+DATA_CRC_SIZE {
		b, err := s.Read()
		if err != nil {
			return bytes, err
		}
		bytes = append(bytes, b)
	}
	return bytes, nil
}
