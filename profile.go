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
}
