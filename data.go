package fitprotocol

import (
	"bufio"
	"errors"
)

const (
	DATA_CRC_SIZE = 2
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
