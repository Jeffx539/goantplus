package messages

import "github.com/jeffx539/goantplus/constants"

// ControlSystemReset resets the antplus dongle
func ControlSystemReset() Message {
	return BuildMessageType(constants.ControlMessageResetSystem, []byte{0x00})
}

// ControlRequest - Refer to 9.5.4.4
func ControlRequest(channel byte, messageID byte) Message {
	return BuildMessageType(constants.ControlMessageRequestMessage, []byte{channel, messageID})
}

// ControlOpenChannel assigns channel type
func ControlOpenChannel(channel byte) Message {
	return BuildMessageType(constants.ControlMessageOpenChannel, []byte{channel})
}

// ControlCloseChannel assigns channel type
func ControlCloseChannel(channel byte) Message {
	return BuildMessageType(constants.ControlMessageCloseChannel, []byte{channel})
}
