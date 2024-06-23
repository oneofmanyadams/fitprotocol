package fitprotocol

// The types ending in "z" just have a different invalid (default) value.
const (
	BASE_TYPE_ENUM    = 0x00
	BASE_TYPE_SINT8   = 0x01
	BASE_TYPE_UINT8   = 0x02
	BASE_TYPE_SINT16  = 0x83
	BASE_TYPE_UINT16  = 0x84
	BASE_TYPE_SINT32  = 0x85
	BASE_TYPE_UINT32  = 0x86
	BASE_TYPE_STRING  = 0x07
	BASE_TYPE_FLOAT32 = 0x88
	BASE_TYPE_FLOAT64 = 0x89
	BASE_TYPE_UINT8Z  = 0x0A
	BASE_TYPE_UINT16Z = 0x8B
	BASE_TYPE_UINT32Z = 0x8C
	BASE_TYPE_BYTE    = 0x0D
	BASE_TYPE_SINT64  = 0x8E
	BASE_TYPE_UINT64  = 0x8F
	BASE_TYPE_UINT64Z = 0x90
)

func BaseTypeName(b byte) string {
	switch b {
	case BASE_TYPE_ENUM:
		return "ENUM"
	case BASE_TYPE_SINT8:
		return "SINT8"
	case BASE_TYPE_UINT8:
		return "UINT8"
	case BASE_TYPE_SINT16:
		return "SINT16"
	case BASE_TYPE_UINT16:
		return "UINT16"
	case BASE_TYPE_SINT32:
		return "SINT32"
	case BASE_TYPE_UINT32:
		return "UINT32"
	case BASE_TYPE_STRING:
		return "STRING"
	case BASE_TYPE_FLOAT32:
		return "FLOAT32"
	case BASE_TYPE_FLOAT64:
		return "FLOAT64"
	case BASE_TYPE_UINT8Z:
		return "UINT8Z"
	case BASE_TYPE_UINT16Z:
		return "UINT16Z"
	case BASE_TYPE_UINT32Z:
		return "UINT32Z"
	case BASE_TYPE_BYTE:
		return "BYTE"
	case BASE_TYPE_SINT64:
		return "SINT64"
	case BASE_TYPE_UINT64:
		return "UINT64"
	case BASE_TYPE_UINT64Z:
		return "UINT64Z"
	}
	return "NONE"
}
