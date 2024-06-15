package fitprotocol

const (
	MSG_RESERVED_SIZE        = 1
	MSG_ARCHITECTURE_SIZE    = 1
	MSG_GLOBAL_NUM_SIZE      = 2
	MSG_FIELD_COUNT_SIZE     = 1
	MSG_FIELD_DEF_SIZE       = 3
	MSG_FIELD_COUNT_SIZE_DEV = 1
	MSG_FIELD_DEF_SIZE_DEV   = 3
)

func MsgFixedContentSize() int {
	return MSG_RESERVED_SIZE +
		MSG_ARCHITECTURE_SIZE +
		MSG_GLOBAL_NUM_SIZE +
		MSG_FIELD_COUNT_SIZE
}

type DefinitionMessage struct {
	Architecture        uint8
	GlobalMessageNumber uint16
	NumberOfFields      uint8
	FieldDefinitions    []FieldDefinition
	DevFlag             bool
	NumberOfDevFields   uint8
	DevFieldDefinitions []FieldDefinition
}

func NewDefinitionMessage(b []byte) DefinitionMessage {
	var dm DefinitionMessage
	dm.Architecture = b[1]
	dm.GlobalMessageNumber = uint16(b[2])<<8 + uint16(b[3])
	dm.NumberOfFields = b[4]
	return dm
}

func (s *DefinitionMessage) AddFieldDef(bytes []byte) {
	var fd FieldDefinition
	for _, b := range bytes {
		fd.Bytes = append(fd.Bytes, b)
	}
	s.FieldDefinitions = append(s.FieldDefinitions, fd)
}

type FieldDefinition struct {
	Bytes []byte
}
