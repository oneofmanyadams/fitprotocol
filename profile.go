package fitprotocol

// Profile information is in Profile.xlsx in SDK
type Profile struct {
	Version      ProfVersion
	CommonFields CommonFields
	Messages     []Message
	Types        []Type
	MesgNum      []MesgNum
}

type ProfVersion struct {
	Major int
	Minor int
	Patch int
	Type  string
}
type CommonFields struct {
	PartIndex    int
	Timestamp    int
	MessageIndex int
}

type Message struct {
	Num         int
	Name        string
	MessagesKey string
	Fields      []Field
}
type Field struct {
	Num           int
	Name          string
	Type          string
	Array         bool
	Scale         int
	Offset        int
	Units         string
	Bits          []int
	Components    []int
	IsAccumulated bool
	HasComponents bool
	Map           []FieldMap
	SubFields     []Field
}

type FieldMap struct {
	Name  string
	Value int
}

type Type struct {
	Name   string
	Values map[int]string
}
type MesgNum struct {
	Name  string
	Index int
}

var PROFILE = Profile{
	Version: ProfVersion{
		Major: 21,
		Minor: 141,
		Patch: 0,
		Type:  "Release"},
	CommonFields: CommonFields{
		PartIndex:    250,
		Timestamp:    253,
		MessageIndex: 254},
	Messages: []Message{
		{Num: 0, Name: "fileId", MessagesKey: "fileIdMesgs",
			Fields: []Field{{
				Num:    0,
				Name:   "type",
				Array:  false,
				Scale:  1,
				Offset: 0,
				Units:  ""},
			},
		},
		{Num: 8, Name: "hrZone", MessagesKey: "hrZoneMesgs",
			Fields: []Field{
				{
					Num:    254,
					Name:   "messageIndex",
					Type:   "messageIndex",
					Scale:  1,
					Offset: 0,
					Units:  "",
				},
				{
					Num:    1,
					Name:   "highBpm",
					Type:   "uint8",
					Scale:  1,
					Offset: 0,
					Units:  "bpm",
				},
				{
					Num:    2,
					Name:   "name",
					Type:   "string",
					Scale:  1,
					Offset: 0,
					Units:  "",
				},
			},
		},
		{Num: 9, Name: "powerZone", MessagesKey: "powerZoneMesgs",
			Fields: []Field{
				{
					Num:    254,
					Name:   "messageIndex",
					Type:   "messageIndex",
					Scale:  1,
					Offset: 0,
					Units:  "",
				},
				{
					Num:    1,
					Name:   "highValue",
					Type:   "uint16",
					Scale:  1,
					Offset: 0,
					Units:  "watts",
				},
				{
					Num:    2,
					Name:   "name",
					Type:   "string",
					Scale:  1,
					Offset: 0,
					Units:  "",
				},
			},
		},
		{Num: 12, Name: "sport", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 18, Name: "session", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 19, Name: "lap", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 20, Name: "record", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 21, Name: "event", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 23, Name: "device_info", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 26, Name: "workout", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 34, Name: "activity", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 206, Name: "field_description", MessagesKey: "",
			Fields: []Field{{}},
		},
		{Num: 207, Name: "developer_data_id", MessagesKey: "",
			Fields: []Field{{}},
		},
		// mesg 65281
		// mesg 65284
		// mesg 65285
	},
	Types:   []Type{},
	MesgNum: []MesgNum{},
}

// Depricate the stuff below, use the js SDK profile as a reference.

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
	{Name: "session", Number: 18}, //what does the SDK do for these?
	{Name: "lap", Number: 19},
	{Name: "record", Number: 20},
	{Name: "event", Number: 21},
	{Name: "device_info", Number: 23},
	{Name: "workout", Number: 26, Fields: []MsgField{
		{Name: "sport", Number: 4},
		{Name: "capabilities", Number: 5},
		{Name: "num_valid_steps", Number: 6},
		{Name: "wkt_name", Number: 8},
		{Name: "sub_sport", Number: 11},
		{Name: "pool_length", Number: 14},
		{Name: "pool_length_unit", Number: 15},
	}},
	{Name: "activity", Number: 34, Fields: []MsgField{
		{Name: "total_timer_time", Number: 0},
		{Name: "num_sessions", Number: 1},
		{Name: "type", Number: 2},
		{Name: "event", Number: 3},
		{Name: "event_type", Number: 4},
		{Name: "local_timestamp", Number: 5},
		{Name: "event_group", Number: 6},
	}},
	// mesg 65281
	// mesg 65284
	// mesg 65285
	{Name: "field_description", Number: 206, Fields: []MsgField{
		{Name: "developer_data_index", Number: 0},
		{Name: "field_definition_number", Number: 1},
		{Name: "fit_base_type_id", Number: 2},
		{Name: "field_name", Number: 3},
		{Name: "array", Number: 4},
		{Name: "components", Number: 5},
		{Name: "scale", Number: 6},
		{Name: "offset", Number: 7},
		{Name: "units", Number: 8},
		{Name: "bits", Number: 9},
		{Name: "accumulate", Number: 10},
		{Name: "fit_base_unit_id", Number: 13},
		{Name: "native_mesg_num", Number: 14},
		{Name: "native_field_num", Number: 15},
	}},
	{Name: "developer_data_id", Number: 207, Fields: []MsgField{
		{Name: "developer_id", Number: 0},
		{Name: "application_id", Number: 1},
		{Name: "manufacturer_id", Number: 2},
		{Name: "developer_data_index", Number: 3},
		{Name: "application_version", Number: 4},
	}},
}
