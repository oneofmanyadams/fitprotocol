package fitprotocol

// Profile information is in Profile.xlsx in SDK
type Profile struct {
}

type MsgNums []MsgNum
type MsgNum struct {
	Name   string
	Number int
}

var MESSAGE_NUMBERS = MsgNums{
	{Name: "file_id", Number: 0},
	{Name: "hr_zone", Number: 8},
	{Name: "power_zone", Number: 9},
	{Name: "sport", Number: 12},
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
