package messages

import (
	"fmt"

	"github.com/jeffx539/goantplus/constants"
)

// Message defined the returntype for the ant+ message, a simple slice of bytes
type Message []byte

// BuildMessage build the final antplus packet provided only with the payload It will return a well formed and packet (Sync, Length,ID,Content) Payload should prepend the MessageID
func BuildMessage(payload []byte) Message {
	var buffer Message
	buffer = append(buffer, constants.AntMessageSync)
	buffer = append(buffer, (byte)(len(payload)-1))
	buffer = append(buffer, payload...)
	buffer = append(buffer, calculateCRC(buffer))
	return buffer
}

// BuildMessageType extends BuildMessage but prepends the ID
func BuildMessageType(messageType byte, payload []byte) Message {
	buffer := make(Message, len(payload)+1)
	buffer[0] = messageType
	copy(buffer[1:], payload)
	return BuildMessage(buffer)
}

func calculateCRC(payload Message) byte {
	var crc byte
	for _, pl := range payload {
		crc ^= pl
	}
	return crc
}

func DebugPrint(payload []byte) {

	for _, elem := range payload {
		fmt.Printf("[ %X ]", elem)
	}

	fmt.Println("")
}
