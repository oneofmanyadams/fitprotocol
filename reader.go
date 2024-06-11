package fitprotocol

import (
	"bufio"
	"os"
)

const (
	HEADER_CRC_SIZE = 2
	DATA_CRC_SIZE   = 2
)

type FitReader struct {
	File      *os.File
	Buffer    *bufio.Reader
	BytesRead int
}

func NewFitReader(file *os.File) (FitReader, error) {
	return FitReader{File: file, Buffer: bufio.NewReader(file)}, nil
}

func (s *FitReader) HeaderSize() (int, error) {
	header_size, read_err := s.PeekBytes(0, 1)
	if read_err != nil {
		return 0, read_err
	}
	return int(header_size[0]), nil
}

func (s *FitReader) HeaderBytes() ([]byte, error) {
	header_size, read_err := s.HeaderSize()
	if read_err != nil {
		return []byte{}, read_err
	}
	return s.PeekBytes(0, header_size)
}

func (s *FitReader) PeekBytes(offset, length int) ([]byte, error) {
	b := make([]byte, length)
	_, read_err := s.File.ReadAt(b, int64(offset))
	if read_err != nil {
		return []byte{}, read_err
	}
	return b, nil
}

// Buffer Methods
func (s *FitReader) ReadByte() (byte, error) {
	s.BytesRead++
	return s.Buffer.ReadByte()
}

func (s *FitReader) ReadBytes(bytes_len int) []byte {
	var return_bytes []byte
	for len(return_bytes) < bytes_len {
		b, err := s.ReadByte()
		if err != nil {
			return return_bytes
		}
		return_bytes = append(return_bytes, b)
	}
	return return_bytes
}

// CRC consolidated checker
func (s *FitReader) CRCs() (bool, bool, error) {
	var header_matches bool
	var data_matches bool
	// Decode header
	header_bytes, err := s.HeaderBytes()
	if err != nil {
		return header_matches, data_matches, err
	}
	header, err := DecodeHeader(header_bytes)
	if err != nil {
		return header_matches, data_matches, err
	}
	// Check Header CRC
	header_crc, err := CalculateCRC(header.HeaderBytes, 0, 11)
	if err != nil {
		return header_matches, data_matches, err
	}
	header_matches = header_crc.Matches(header.CRC)

	// Calaculate data CRC
	data_bytes, err := s.PeekBytes(header.HeaderSize, int(header.DataSize))
	if err != nil {
		return header_matches, data_matches, err
	}
	data_crc, err := CalculateCRC(data_bytes, 0, int(header.DataSize)-1)
	if err != nil {
		return header_matches, data_matches, err
	}
	// Get provided Data CRC from end of file
	data_crc_bytes, err := s.PeekBytes(header.TotalSize(), DATA_CRC_SIZE)
	if err != nil {
		return header_matches, data_matches, err
	}
	// compare data crcs.
	data_matches = data_crc.Matches(data_crc_bytes)
	return header_matches, data_matches, nil
}
