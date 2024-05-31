package fitprotocol

import (
	"bufio"
	"fmt"
)

const (
	DATA_CRC_SIZE = 2
)

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
		ReadLines:  0,
		Position:   0,
		ReadLimit:  data_size}
}

func (s *Data) Read() byte {
	b, err := s.Reader.ReadByte()
	if err != nil {
		fmt.Println(err)
	}
	s.ReadLines++
	s.Position++
	return b
}

func (s *Data) ReadRemaining() []byte {
	var bytes []byte
	for s.ReadLines < s.DataSize+DATA_CRC_SIZE {
		bytes = append(bytes, s.Read())
	}
	return bytes
}
