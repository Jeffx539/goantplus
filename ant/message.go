package ant

import (
	"encoding/binary"
	"fmt"
)

// Message defined the returntype for the ant+ message, a simple slice of bytes
type Message []byte

// MessageType defines the AntPlus data structure
type MessageType struct {
	sync     byte
	id       byte
	length   byte
	payload  []byte
	checksum byte
}

// BuildMessage build the final antplus packet provided only with the payload It will return a well formed and packet (Sync, Length,ID,Content) Payload should prepend the MessageID
func BuildMessage(payload []byte) Message {
	var buffer Message
	buffer = append(buffer, AntMessageSync)
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

// MessageSystemReset resets the antplus dongle
func MessageSystemReset() Message {
	return BuildMessageType(AntControlMessageResetSystem, []byte{0x00})
}

// MessageRequest - Refer to 9.5.4.4
func MessageRequest(channel byte, messageID byte) Message {
	return BuildMessageType(AntControlMessageRequest, []byte{channel, messageID})
}

// MessageSetNetworkKey - Refer to
func MessageSetNetworkKey(channel byte, key []byte) Message {
	return BuildMessageType(AntMessageSetNetworkKey, append([]byte{byte(channel)}, key...))
}

// MessageAssignChannel assigns channel type
func MessageAssignChannel(channel byte, typ byte) Message {
	return BuildMessageType(AntMessageAssignChannel, []byte{channel, typ, 0x00})
}

// MessageUnassignChannel assigns channel type
func MessageUnassignChannel(channel byte, typ byte) Message {
	return BuildMessageType(AntMessageUnassignChannel, []byte{channel, typ})
}

// MessageSetChannelID sets channel ID
func MessageSetChannelID(channel byte, deviceNo byte, deviceTyp1 byte, deviceTyp2 byte, transmissionTyp byte) Message {
	return BuildMessageType(AntMessageSetChannelID, []byte{channel, deviceNo, deviceTyp1, deviceTyp2, transmissionTyp})
}

// MessageSetChannelRFFrequency sets channel ID
func MessageSetChannelRFFrequency(channel byte, freq byte) Message {
	return BuildMessageType(AntMessageSetChannelRFFrequency, []byte{channel, freq})
}

// MessageSetChannelRFFrequency sets channel ID
func MessageSetChannelPeriod(channel byte, period uint16) Message {

	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, period)
	return BuildMessageType(AntMessageSetChannelPeriod, []byte{channel, b[0], b[1]})
}

// MessageOpenChannel assigns channel type
func MessageOpenChannel(channel byte) Message {
	return BuildMessageType(AntMessageOpenChannel, []byte{channel})
}

// MessageOpenRXScanMode assigns channel type
func MessageOpenRXScanMode(channel byte) Message {
	return BuildMessageType(AntOpenRXScanMode, []byte{channel, 0x00})
}

// DecodeMessage Marshals the payload into the struct
func (msg MessageType) DecodeMessage(payload []byte) {
	msg.sync = payload[0]
	msg.id = payload[1]
	msg.length = payload[2]
	msg.payload = payload[3 : msg.length-1]
	msg.checksum = payload[msg.length]
}

func DebugPrint(payload []byte) {

	if payload[0] == 0x4e {
		fmt.Println("[HR] ", payload[9])
		return
	}

	for _, elem := range payload {
		fmt.Printf("[ %X ]", elem)
	}

	fmt.Println("")
}
