package pkg

import (
	"errors"
	"io"

	"github.com/jeffx539/goantplus/pkg/constants"
	"github.com/jeffx539/goantplus/pkg/messages"

	"github.com/jacobsa/go-serial/serial"
)

type Device struct {
	stream   io.ReadWriteCloser
	channels []*Channel
}

func (d Device) GetNextAvailableChannel() int {
	return len(d.channels)
}

// CreateSerialDevice Initialises serial port with correct parameters
func CreateSerialDevice(path string) (io.ReadWriteCloser, error) {

	options := serial.OpenOptions{
		PortName:        path,
		BaudRate:        4800,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)

	return port, err
}

// MakeDevice Initalises and Device Structure
func MakeDevice(path string) *Device {

	dev, err := CreateSerialDevice(path)
	if err != nil {
		return nil
	}

	return &Device{
		stream: dev,
	}

}

// InitialiseDevice resets system and allocates
func (d *Device) InitialiseDevice() error {
	_, err := d.stream.Write(messages.ControlSystemReset())

	if err != nil {
		return err
	}

	return nil
}

// IntialiseChannel resets system and allocates
func (d *Device) IntialiseChannel(channel byte) *Channel {
	c := MakeChannel(channel, d.stream)
	d.channels = append(d.channels, c)
	return c
}

// Loop a
func (d *Device) Loop() {

	for {
		payload, _ := parseMessage(d.stream)
		if payload[0] == 78 {
			d.channels[payload[1]].EmitEvent(payload)
		}

	}
}

func parseMessage(port io.ReadWriteCloser) ([]byte, error) {

	data := make([]byte, 1)
	port.Read(data)

	if data[0] != constants.AntMessageSync {
		return nil, errors.New("Incorrect Sync")
	}

	port.Read(data)

	length := data[0]
	data = make([]byte, length+2)
	port.Read(data)

	return data[0 : length+1], nil

}
