package fitprotocol

import "fmt"

// Create a 2nd helper func for reading data messages.

// Third func for reading timestamp messages?

func ReadDefMsg(fit_reader *FitReader) DefinitionMessage {
	fmt.Println("----------Definition Message---------------------")
	// Read Message Type
	msg_header_bytes, _ := fit_reader.ReadByte()
	fmt.Println("")
	fmt.Println("--Message Header:")
	msg_header := ParseMessageHeader(msg_header_bytes)
	fmt.Printf("%+v\n", msg_header)
	// read fixed length part of def message
	fixed_msg_bytes, _ := fit_reader.ReadBytes(MsgFixedContentSize())
	def_msg := NewDefinitionMessage(fixed_msg_bytes)
	def_msg.DevFlag = msg_header.DevFlag
	fmt.Println("")
	fmt.Println("--Definition message:")
	fmt.Printf("%+v\n", def_msg)
	// read variable length part of data message
	for len(def_msg.FieldDefinitions) < int(def_msg.NumberOfFields) {
		field_bytes, _ := fit_reader.ReadBytes(MSG_FIELD_DEF_SIZE)
		def_msg.AddFieldDef(field_bytes)
	}
	// Read dev data (if applicable)
	if def_msg.DevFlag {
		dev_field_number, _ := fit_reader.ReadByte()
		def_msg.NumberOfDevFields = uint8(dev_field_number)
		fmt.Println("Number of dev fields: ", def_msg.NumberOfDevFields)
		for len(def_msg.DevFieldDefinitions) < int(def_msg.NumberOfDevFields) {
			dev_f_bytes, _ := fit_reader.ReadBytes(MSG_FIELD_DEF_SIZE_DEV)
			def_msg.AddDevFieldDef(dev_f_bytes)
		}
	}
	// display field definition data
	for _, field := range def_msg.FieldDefinitions {
		data_type, err := BASE_TYPES.GetBaseType(field.BaseType)
		/*
			No matching type found
			Where are these missing types?
		*/
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", data_type)
	}
	fmt.Println("")
	// display dev definition data
	if def_msg.DevFlag {
		for _, dev_field := range def_msg.DevFieldDefinitions {
			fmt.Printf("%+v\n", dev_field)
		}
	}
	return def_msg
}

func ReadDataMsg(fit_reader *FitReader, def_msg DefinitionMessage) {
	fmt.Println("-------------Data Message------------------------")
	// Read data message header
	fmt.Println("")
	fmt.Println("--Message Header:")
	msg_header_bytes, _ := fit_reader.ReadByte()
	fmt.Printf("\n%+v\n", ParseMessageHeader(msg_header_bytes))
	// chunk out data message based on def message
	fmt.Println("")
	fmt.Println("--Data message:")
	data_bytes, _ := fit_reader.ReadBytes(def_msg.DataMessageSize())
	fmt.Println(def_msg.ParseDataMessage(data_bytes))
	fmt.Println("")
}

func ReadTimestampMsg(fit_reader *FitReader, def_msg DefinitionMessage) {
	fmt.Println("-------------Data Message------------------------")
	// Read data message header
	fmt.Println("")
	fmt.Println("--Message Header:")
	msg_header_bytes, _ := fit_reader.ReadByte()
	fmt.Printf("\n%+v\n", ParseMessageHeader(msg_header_bytes))
}
