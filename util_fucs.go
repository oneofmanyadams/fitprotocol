package fitprotocol

import "fmt"

// Change this to be just for reading definition messages.

// Create a 2nd helper func for reading data messages.

// Third func for reading timestamp messages?

func ReadDefMsg(fit_reader *FitReader) DefinitionMessage {
	fmt.Println("-------------------------------------------------")
	// Read Message Type
	msg_header_bytes, _ := fit_reader.ReadByte()
	fmt.Println("")
	fmt.Println("--Message Header:")
	fmt.Printf("%+v\n", ParseMessageHeader(msg_header_bytes))
	// read fixed length part of def message
	fixed_msg_bytes, _ := fit_reader.ReadBytes(MsgFixedContentSize())
	def_msg := NewDefinitionMessage(fixed_msg_bytes)
	fmt.Println("")
	fmt.Println("--Definition message:")
	fmt.Printf("%+v\n", def_msg)
	// read variable length part of data message
	for len(def_msg.FieldDefinitions) < int(def_msg.NumberOfFields) {
		field_bytes, _ := fit_reader.ReadBytes(MSG_FIELD_DEF_SIZE)
		def_msg.AddFieldDef(field_bytes)
	}
	// display field definition data
	for _, field := range def_msg.FieldDefinitions {
		data_type, err := BASE_TYPES.GetBaseType(field.BaseType)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", data_type)
	}
	fmt.Println("")
	return def_msg
}
