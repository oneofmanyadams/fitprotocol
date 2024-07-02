package fitprotocol

import "fmt"

func ReadRecords(fit_reader *FitReader) {
	// Read Message Type
	msg_header_bytes, _ := fit_reader.ReadByte()
	fmt.Println("Message Header:")
	fmt.Printf("%+v\n", ParseMessageHeader(msg_header_bytes))
	// read fixed length part of def message
	fixed_msg_bytes, _ := fit_reader.ReadBytes(MsgFixedContentSize())
	def_msg := NewDefinitionMessage(fixed_msg_bytes)
	fmt.Println("Definition message:")
	fmt.Printf("%+v\n", def_msg)
}
