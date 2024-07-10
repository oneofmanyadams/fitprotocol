package fitprotocol

type TimestampDataMessage struct {
	Header           int // 7th bit
	LocalMessageType int // bits 5 and 6
	TimeOffset       int // bits 0 through 4
}
