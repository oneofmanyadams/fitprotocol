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

func NewDefinitionMessage(b []byte) DefinitionMessage {
	var dm DefinitionMessage
	return dm
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

func (s *DefinitionMessage) AddFieldDef() {}

type FieldDefinition struct {
}
