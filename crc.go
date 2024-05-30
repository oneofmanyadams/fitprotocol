package fitprotocol

import (
	"encoding/binary"
	"errors"
)

var CRC_BUFFER_BOUNDS_ERROR = errors.New("Attempted to read outside of buffer.")

var CRC_TABLE = []uint16{
	0x0000, 0xCC01, 0xD801, 0x1400,
	0xF001, 0x3C00, 0x2800, 0xE401,
	0xA001, 0x6C00, 0x7800, 0xB401,
	0x5000, 0x9C01, 0x8801, 0x4400,
}

type CRC struct {
	crc        uint16
	bytes_seen int
}

func CalculateCRC(buffer []byte, start, end int) (CRC, error) {
	var crc CRC
	err := crc.AddBytes(buffer, start, end)
	return crc, err
}

func (s *CRC) CRC() uint16 {
	return s.crc
}

func (s *CRC) Matches(want_crc []byte) bool {
	return s.crc == binary.LittleEndian.Uint16(want_crc)
}

func (s *CRC) updateCRC(b byte) {
	var tmp uint16

	// compute checksum of lower four bits of byte.
	tmp = CRC_TABLE[s.crc&0xF]
	s.crc = (s.crc >> 4) & 0x0FFF
	s.crc = s.crc ^ tmp ^ CRC_TABLE[b&0xF]

	// compute checksum of upper four bits of byte
	tmp = CRC_TABLE[s.crc&0xF]
	s.crc = (s.crc >> 4) & 0x0FFF
	s.crc = s.crc ^ tmp ^ CRC_TABLE[(b>>4)&0xF]
}

func (s *CRC) AddBytes(buff []byte, start, end int) error {
	if start < 0 || end >= len(buff) {
		return CRC_BUFFER_BOUNDS_ERROR
	}
	for i := start; i <= end; i++ {
		s.updateCRC(buff[i])
		s.bytes_seen++
	}
	return nil
}
