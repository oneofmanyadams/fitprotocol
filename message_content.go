package fitprotocol

import (
	"encoding/binary"
	"errors"
	"strconv"
)

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
	Architecture        uint8 // 0 is little endian.
	GlobalMessageNumber uint16
	NumberOfFields      uint8
	FieldDefinitions    []FieldDefinition
	DevFlag             bool
	NumberOfDevFields   uint8
	DevFieldDefinitions []DevFieldDefinition
}

func NewDefinitionMessage(b []byte) DefinitionMessage {
	var dm DefinitionMessage
	dm.Architecture = b[1]
	//dm.GlobalMessageNumber = uint16(b[2])<<8 + uint16(b[3])
	if dm.Architecture == 0 {
		dm.GlobalMessageNumber = binary.LittleEndian.Uint16([]byte{b[2], b[3]})
	} else {
		dm.GlobalMessageNumber = binary.BigEndian.Uint16([]byte{b[2], b[3]})
	}
	dm.NumberOfFields = b[4]
	return dm
}

func (s *DefinitionMessage) AddFieldDef(bytes []byte) error {
	var fd FieldDefinition
	if len(bytes) != MSG_FIELD_DEF_SIZE {
		return errors.New("Incorrect amount of bytes provided.")
	}
	fd.Number = bytes[0]
	fd.Size = bytes[1]
	fd.BaseType = bytes[2]
	fd.Bytes = bytes
	s.FieldDefinitions = append(s.FieldDefinitions, fd)
	return nil
}

func (s *DefinitionMessage) DataMessageSize() int {
	var size int
	for _, f := range s.FieldDefinitions {
		size = size + int(f.Size)
	}
	return size
}

func (s *DefinitionMessage) ParseDataMessage(b []byte) ([]DataPoint, error) {
	var datas []DataPoint
	var bytes_read int
	if len(b) != s.DataMessageSize() {
		return datas, errors.New("Size of provided bytes does not match expected data fields size.")
	}
	for len(datas) < len(s.FieldDefinitions) {
		var dta DataPoint
		size := s.FieldDefinitions[len(datas)].Size
		type_bytes := s.FieldDefinitions[len(datas)].BaseType
		dta.Bytes = b[bytes_read : bytes_read+int(size)]
		data_type, dt_err := BASE_TYPES.GetBaseType(type_bytes)
		if dt_err != nil {
			return []DataPoint{}, dt_err
		}
		var conv_err error
		dta.Type, conv_err = data_type.ConvertData(dta.Bytes, s.Architecture)
		if conv_err != nil {
			return []DataPoint{}, errors.Join(conv_err, errors.New(data_type.Name))
		}

		datas = append(datas, dta)
		bytes_read = bytes_read + int(size)
	}
	return datas, nil
}

func (s *DefinitionMessage) MessageNumber() (MsgNum, error) {
	for _, msg_num := range MESSAGE_NUMBERS {
		if msg_num.Number == int(s.GlobalMessageNumber) {
			return msg_num, nil
		}
	}
	return MsgNum{}, errors.New("NO MATCHING MESSAGE NAME FOR " + strconv.Itoa(int(s.GlobalMessageNumber)))
}

func (s *DefinitionMessage) MessageName() string {
	msg_num, err := s.MessageNumber()
	if err != nil {
		return err.Error()
	}
	return msg_num.Name
}

type DataPoint struct {
	Bytes []byte
	Type  string
}

type FieldDefinition struct {
	Number   uint8 // Defined in the global fit profile for the specified message.
	Size     uint8 // size in bytes of specifieed fit message's field.
	BaseType uint8 // (unsigned char, etc...) defined in fit.h in SDK
	Bytes    []byte
}

func (s *FieldDefinition) FieldName(msg_num MsgNum) string {
	for _, field := range msg_num.Fields {
		if field.Number == s.Number {
			return field.Name
		}
	}
	return "NO MATCHING FIELD NAME FOR " + msg_num.Name + " " + strconv.Itoa(int(s.Number))
}

type DevFieldDefinition struct {
	Number       uint8 // maps to field_def_number of a field_Description msg.
	Size         uint8 // size in bytes of specifieed fit message's field.
	DevDataIndex uint8 // maps to developer_data_index of a developer_data_id msg
	Bytes        []byte
}

func (s *DefinitionMessage) AddDevFieldDef(bytes []byte) error {
	var fd DevFieldDefinition
	if len(bytes) != MSG_FIELD_DEF_SIZE_DEV {
		return errors.New("Incorrect amount of bytes provided.")
	}
	fd.Number = bytes[0]
	fd.Size = bytes[1]
	fd.DevDataIndex = bytes[2]
	fd.Bytes = bytes
	s.DevFieldDefinitions = append(s.DevFieldDefinitions, fd)
	return nil
}

func (s *DefinitionMessage) ParseDevMessage(b []byte) {

}

func (s *DefinitionMessage) ParseTimestampMessage(b []byte) {

}

type TimestampMessage struct {
	HeaderFlag       int // 7th bit
	LocalMessageType int // bits 5 and 6
	TimeOffset       int // bits 0 through 4
}
