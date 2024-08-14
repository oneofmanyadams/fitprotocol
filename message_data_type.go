package fitprotocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type DataTypes []DataType

func (s DataTypes) GetBaseType(btf uint8) (DataType, error) {
	for _, data_type := range s {
		if data_type.BaseTypeField == btf {
			return data_type, nil
		}
	}
	return DataType{}, errors.New("No matching type found")
}

type DataType struct {
	BaseTypeNumber int
	EndianAbility  bool
	BaseTypeField  uint8
	Name           string
	InvalidValue   uint64
	Size           uint8
}

func (s *DataType) ConvertData(b []byte, arch uint8) (string, error) {
	switch s.Name {
	case "enum":
		return "an enum", nil
	case "sint8":
		return fmt.Sprint(int8(b[0])), nil
	case "uint8":
		return fmt.Sprint(uint8(b[0])), nil
	case "sint16":
		if arch == 0 {
			return fmt.Sprint(int16(binary.LittleEndian.Uint16(b))), nil
		} else {
			return fmt.Sprint(int16(binary.BigEndian.Uint16(b))), nil
		}
	case "uint16":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint16(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint16(b)), nil
		}
	case "sint32":
		if arch == 0 {
			return fmt.Sprint(int32(binary.LittleEndian.Uint32(b))), nil
		} else {
			return fmt.Sprint(int32(binary.BigEndian.Uint32(b))), nil
		}
	case "uint32":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint32(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint32(b)), nil
		}
	case "string":
		return fmt.Sprintf("%s", b), nil //string(b), nil
	case "float32":
		var f float32
		if arch == 0 {
			binary.Read(bytes.NewReader(b), binary.LittleEndian, &f)
		} else {
			binary.Read(bytes.NewReader(b), binary.BigEndian, &f)
		}
		return fmt.Sprint(f), nil
	case "float64":
		var f float64
		if arch == 0 {
			binary.Read(bytes.NewReader(b), binary.LittleEndian, &f)
		} else {
			binary.Read(bytes.NewReader(b), binary.BigEndian, &f)
		}
		return fmt.Sprint(f), nil
	case "uint8z":
		return fmt.Sprint(uint8(b[0])), nil
	case "uint16z":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint16(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint16(b)), nil
		}
	case "uint32z":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint32(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint32(b)), nil
		}
	case "byte":
		return fmt.Sprintf("%08b", b[0]), nil
	case "sint64":
		if arch == 0 {
			return fmt.Sprint(int64(binary.LittleEndian.Uint64(b))), nil
		} else {
			return fmt.Sprint(int64(binary.BigEndian.Uint64(b))), nil
		}
	case "uint64":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint64(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint64(b)), nil
		}
	case "uint64z":
		if arch == 0 {
			return fmt.Sprint(binary.LittleEndian.Uint64(b)), nil
		} else {
			return fmt.Sprint(binary.BigEndian.Uint64(b)), nil
		}
	}
	return "", errors.New("No matching type to covnert.")
}

var BASE_TYPES = DataTypes{
	DataType{
		BaseTypeNumber: 0,
		EndianAbility:  false,
		BaseTypeField:  0x00,
		Name:           "enum",
		InvalidValue:   0xFF,
		Size:           1},
	DataType{
		BaseTypeNumber: 1,
		EndianAbility:  false,
		BaseTypeField:  0x01,
		Name:           "sint8",
		InvalidValue:   0x7F,
		Size:           1},
	DataType{
		BaseTypeNumber: 2,
		EndianAbility:  false,
		BaseTypeField:  0x02,
		Name:           "uint8",
		InvalidValue:   0xFF,
		Size:           1},
	DataType{
		BaseTypeNumber: 3,
		EndianAbility:  true,
		BaseTypeField:  0x83,
		Name:           "sint16",
		InvalidValue:   0x7FFF,
		Size:           2},
	DataType{
		BaseTypeNumber: 4,
		EndianAbility:  true,
		BaseTypeField:  0x84,
		Name:           "uint16",
		InvalidValue:   0xFFFF,
		Size:           2},
	DataType{
		BaseTypeNumber: 5,
		EndianAbility:  true,
		BaseTypeField:  0x85,
		Name:           "sint32",
		InvalidValue:   0x7FFFFFFF,
		Size:           4},
	DataType{
		BaseTypeNumber: 6,
		EndianAbility:  true,
		BaseTypeField:  0x86,
		Name:           "uint32",
		InvalidValue:   0xFFFFFFFF,
		Size:           4},
	DataType{
		BaseTypeNumber: 7,
		EndianAbility:  false,
		BaseTypeField:  0x07,
		Name:           "string",
		InvalidValue:   0x00,
		Size:           1},
	DataType{
		BaseTypeNumber: 8,
		EndianAbility:  true,
		BaseTypeField:  0x88,
		Name:           "float32",
		InvalidValue:   0xFFFFFFFF,
		Size:           4},
	DataType{
		BaseTypeNumber: 9,
		EndianAbility:  true,
		BaseTypeField:  0x89,
		Name:           "float64",
		InvalidValue:   0xFFFFFFFFFFFFFFFF,
		Size:           8},
	DataType{
		BaseTypeNumber: 10,
		EndianAbility:  false,
		BaseTypeField:  0x0A,
		Name:           "uint8z",
		InvalidValue:   0x00,
		Size:           1},
	DataType{
		BaseTypeNumber: 11,
		EndianAbility:  true,
		BaseTypeField:  0x8B,
		Name:           "uint16z",
		InvalidValue:   0x0000,
		Size:           2},
	DataType{
		BaseTypeNumber: 12,
		EndianAbility:  true,
		BaseTypeField:  0x8C,
		Name:           "uint32z",
		InvalidValue:   0x00000000,
		Size:           4},
	DataType{
		BaseTypeNumber: 13,
		EndianAbility:  false,
		BaseTypeField:  0x0D,
		Name:           "byte",
		InvalidValue:   0xFF,
		Size:           1},
	DataType{
		BaseTypeNumber: 14,
		EndianAbility:  true,
		BaseTypeField:  0x8E,
		Name:           "sint64",
		InvalidValue:   0x7FFFFFFFFFFFFFFF,
		Size:           8},
	DataType{
		BaseTypeNumber: 15,
		EndianAbility:  true,
		BaseTypeField:  0x8F,
		Name:           "uint64",
		InvalidValue:   0xFFFFFFFFFFFFFFFF,
		Size:           8},
	DataType{
		BaseTypeNumber: 16,
		EndianAbility:  true,
		BaseTypeField:  0x90,
		Name:           "uint64z",
		InvalidValue:   0x0000000000000000,
		Size:           8},
}
