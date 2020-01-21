package messages

import (
	"encoding/binary"

	"github.com/jeffx539/goantplus/constants"
)

// ConfigurationSetNetworkKey - Refer to
func ConfigurationSetNetworkKey(channel byte, key []byte) Message {
	return BuildMessageType(constants.ConfigurationMessageSetNetworkKey, append([]byte{byte(channel)}, key...))
}

// ConfigurationAssignChannel assigns channel type
func ConfigurationAssignChannel(channel byte, typ byte) Message {
	return BuildMessageType(constants.ConfigurationMessageAssignChannel, []byte{channel, typ, 0x00})
}

// ConfigurationUnassignChannel assigns channel type
func ConfigurationUnassignChannel(channel byte, typ byte) Message {
	return BuildMessageType(constants.ConfigurationMessageUnassignChannel, []byte{channel, typ})
}

// ConfigurationSetChannelID sets channel ID
func ConfigurationSetChannelID(channel byte, deviceNo byte, deviceTyp1 byte, deviceTyp2 byte, transmissionTyp byte) Message {
	return BuildMessageType(constants.ConfigurationMessageSetChannelID, []byte{channel, deviceNo, deviceTyp1, deviceTyp2, transmissionTyp})
}

// ConfigurationSetChannelRFFrequency sets channel ID
func ConfigurationSetChannelRFFrequency(channel byte, freq byte) Message {
	return BuildMessageType(constants.ConfigurationMessageSetChannelRFFreq, []byte{channel, freq})
}

// ConfigurationSetChannelPeriod sets channel Period
func ConfigurationSetChannelPeriod(channel byte, period uint16) Message {

	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, period)
	return BuildMessageType(constants.ConfigurationMessageSetChannelPeriod, []byte{channel, b[0], b[1]})
}

// ConfigurationOpenChannel assigns channel type
func ConfigurationOpenChannel(channel byte) Message {
	return BuildMessageType(constants.ControlMessageOpenChannel, []byte{channel})
}

// ConfigurationOpenRXScanMode assigns channel type
func ConfigurationOpenRXScanMode(channel byte) Message {
	return BuildMessageType(constants.ConfigurationMessageOpenRxScanMode, []byte{channel, 0x00})
}
