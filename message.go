package fitprotocol

const (
	TIME_MSG_MASK       = 0x80
	DEF_MSG_MASK        = 0x40
	DATA_MSG_MASK       = 0x00
	LOCAL_MSG_MASK      = 0xF0 // 0x0F
	TIME_LOCAL_MSG_MASK = 0x9F
	DEV_DATA_MASK       = 0x20
)

type MessageHeader struct {
	Byte         uint8
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
	mh.Byte = header_byte
	SetMessageHeaderType(&mh)
	SetLocalMessageType(&mh)
	return mh
}

func SetMessageHeaderType(header *MessageHeader) {
	switch {
	case header.Byte&TIME_MSG_MASK == TIME_MSG_MASK:
		header.IsTimestamp = true
	case header.Byte&DEF_MSG_MASK == DEF_MSG_MASK:
		header.IsDef = true
	case header.Byte&DATA_MSG_MASK == DATA_MSG_MASK:
		header.IsData = true
	}
}

func SetLocalMessageType(header *MessageHeader) {
	switch {
	case header.IsTimestamp:
		header.LocalMsgType = header.Byte &^ TIME_LOCAL_MSG_MASK
	case header.IsDef:
		header.LocalMsgType = header.Byte &^ LOCAL_MSG_MASK
	case header.IsData:
		header.LocalMsgType = header.Byte &^ LOCAL_MSG_MASK
	}
}
