package fitprotocol

const (
	TIME_MSG_MASK      = 0x80
	DEF_MSG_MASK       = 0x40
	DATA_MSG_MASK      = 0x00
	LOCAL_MSG_NUM_MASK = 0x0F
	DEV_DATA_MASK      = 0x20
)

const (
	TS_MESSAGE   = "timestamp_message"
	DEF_MESSAGE  = "def_message"
	DATA_MESSAGE = "data_message"
)

type MessageHeader struct {
	IsTimestamp  bool
	IsHeader     bool
	IsData       bool
	MsgTypeSpec  bool
	LocalMsgType int16
}

func MessageType(msg_header_byte byte) string {
	switch {
	case msg_header_byte&TIME_MSG_MASK == TIME_MSG_MASK:
		return TS_MESSAGE
	case msg_header_byte&DEF_MSG_MASK == DEF_MSG_MASK:
		return DEF_MESSAGE
	case msg_header_byte&DATA_MSG_MASK == DATA_MSG_MASK:
		return DATA_MESSAGE
	}
	return "NONE"
}
