package fitprotocol

import (
	"errors"
)

type Header struct {
	Size            int
	ProtocolVersion string
	ProfileVersion  uint16
	DataSize        uint32
	DataType        string
	CRC             int
}

func HeaderParts(h []byte) (size, proto, prof, dsize, dtype, crc []byte, err error) {
	if len(h) != 12 && len(h) != 14 {
		err = errors.New("invalid fit header")
		return
	}
	size = h[0:1]
	proto = h[1:2]
	prof = h[2:4]
	dsize = h[4:8]
	dtype = h[8:12]
	if len(h) == 14 {
		crc = h[12:13]
	}
	return
}
