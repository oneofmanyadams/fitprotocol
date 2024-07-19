package fitprotocol

// Profile information is in Profile.xlsx in SDK
type Profile struct {
}

type MsgNums []MsgNum
type MsgNum struct {
	Name   string
	Number int
	Fields []MsgField
}
type MsgField struct {
	Name   string
	Number uint8
}

var MESSAGE_NUMBERS = MsgNums{
	{Name: "file_id", Number: 0, Fields: []MsgField{
		{Name: "type", Number: 0},
		{Name: "manufacturer", Number: 1},
		{Name: "product", Number: 2},
		{Name: "serial number", Number: 3},
		{Name: "time created", Number: 4},
		{Name: "number", Number: 5},
		{Name: "product name", Number: 8},
	}},
	{Name: "hr_zone", Number: 8, Fields: []MsgField{
		{Name: "message_index", Number: 254},
		{Name: "high_bpm", Number: 1},
		{Name: "name", Number: 2},
	}},
	{Name: "power_zone", Number: 9, Fields: []MsgField{
		{Name: "message_index", Number: 254},
		{Name: "high_value", Number: 1},
		{Name: "name", Number: 2},
	}},
	{Name: "sport", Number: 12, Fields: []MsgField{
		{Name: "sport", Number: 0},
		{Name: "sub_sport", Number: 1},
		{Name: "name", Number: 3},
	}},
	{Name: "session", Number: 18},
	{Name: "lap", Number: 19},
	{Name: "record", Number: 20},
	{Name: "event", Number: 21},
	{Name: "device_info", Number: 23},
	{Name: "workout", Number: 26},
	{Name: "activity", Number: 34},
	// mesg 65281
	// mesg 65284
	// mesg 65285
	{Name: "field_description", Number: 206},
	{Name: "developer_data_id", Number: 207},
}
