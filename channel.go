package main

import (
	"io"

	"github.com/jeffx539/goantplus/constants"
	"github.com/jeffx539/goantplus/messages"
)

type ChannelMessageCallback func(payload []byte)

// Channel Struct
type Channel struct {
	ID        uint8
	stream    io.ReadWriteCloser
	callbacks []ChannelMessageCallback

	// Channel Profile Interface ?
}

// MakeChannel Create a Channel with a specified channelID
func MakeChannel(channID uint8, stream io.ReadWriteCloser) *Channel {
	return &Channel{ID: channID, stream: stream}
}

// AssignChannel Create a Channel with a specified channelID
func (c Channel) AssignChannel(typ byte) {
	c.stream.Write(messages.BuildMessageType(constants.ConfigurationMessageAssignChannel, []byte{c.ID, typ, 0x00}))
}

// SetChannelID a
func (c Channel) SetChannelID(deviceNo byte, deviceTyp1 byte, deviceTyp2 byte, transmissionTyp byte) {
	c.stream.Write(messages.ConfigurationSetChannelID(c.ID, deviceNo, deviceTyp1, deviceTyp2, transmissionTyp))
}

// SetChannelPeriod Sets channel period
func (c Channel) SetChannelPeriod(period uint16) {
	c.stream.Write(messages.ConfigurationSetChannelPeriod(c.ID, period))
}

// SetNetworkKey a
func (c Channel) SetNetworkKey(key []byte) {
	c.stream.Write(messages.ConfigurationSetNetworkKey(c.ID, key))
}

// Open Opens Channel
func (c Channel) Open() {
	c.stream.Write(messages.ControlOpenChannel(c.ID))
}

// Close Closes Channel
func (c Channel) Close() {
	c.stream.Write(messages.ControlCloseChannel(c.ID))
}

// SetRFFrequency a
func (c Channel) SetRFFrequency(freq byte) {
	c.stream.Write(messages.ConfigurationSetChannelRFFrequency(c.ID, freq))
}

// EmitEvent call event on all registered functions
func (c *Channel) EmitEvent(pl []byte) {
	for _, cb := range c.callbacks {
		cb(pl)
	}
}

// RegisterCallback - Registers a callback
func (c *Channel) RegisterCallback(cb ChannelMessageCallback) {
	c.callbacks = append(c.callbacks, cb)
}
