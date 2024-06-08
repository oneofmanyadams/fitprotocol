package fitprotocol

const (
	TIME_MSG_MASK      = 0x80
	DEF_MSG_MASK       = 0x40
	DATA_MSG_MASK      = 0x00
	LOCAL_MSG_NUM_MASK = 0x0F
	DEV_DATA_MASK      = 0x20
)

const (
	NOT_A_MESSAGE = 0
	TIME_MESSAGE  = 1
	DEF_MESSAGE   = 2
	DATA_MESSAGE  = 3
)

type MessageHeader struct {
	HeaderByte   uint8
	IsTimestamp  bool
	IsDef        bool
	IsData       bool
	DevFlag      bool
	MsgTypeSpec  bool
	LocalMsgType uint8
	TimeOffset   uint32
}

func ParseMessageHeader(header_byte byte) MessageHeader {
	var mh MessageHeader
	mh.HeaderByte = header_byte
	MessageType(&mh)
	mh.LocalMsgType = header_byte &^ 0xF0
	return mh
}

func MessageType(header *MessageHeader) {
	switch {
	case header.HeaderByte&TIME_MSG_MASK == TIME_MSG_MASK:
		header.IsTimestamp = true
	case header.HeaderByte&DEF_MSG_MASK == DEF_MSG_MASK:
		header.IsDef = true
	case header.HeaderByte&DATA_MSG_MASK == DATA_MSG_MASK:
		header.IsData = true
	}
}
