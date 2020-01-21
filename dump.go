package main

import (
	"fmt"
	"io"
	"log"

	"github.com/jacobsa/go-serial/serial"
	"github.com/jeffx539/goantplus/constants"
	"github.com/jeffx539/goantplus/messages"
)

func parseMessage(port io.ReadWriteCloser) {

	data := make([]byte, 1)
	port.Read(data)

	if data[0] != constants.AntMessageSync {
		panic("Incorrect Sync")
	}

	port.Read(data)

	length := data[0]
	data = make([]byte, length+2)
	port.Read(data)

	messages.DebugPrint(data)

}

func main() {
	fmt.Println("Hello World!")

	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        4800,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()
	port.Write(messages.ControlSystemReset())
	parseMessage(port)

	// port.Write(ant.MessageRequest(0, ant.AntMessageParamCapability))
	// parseMessage(port)

	port.Write(messages.ConfigurationSetNetworkKey(0, ([]byte)(constants.ANTNetworkKey)))
	parseMessage(port)

	port.Write(messages.ConfigurationAssignChannel(0, 0x40))
	parseMessage(port)

	port.Write(messages.ConfigurationSetChannelID(0, 0, 0, 0, 0))
	parseMessage(port)

	port.Write(messages.ConfigurationSetChannelRFFrequency(0, 0x39))
	parseMessage(port)

	port.Write(messages.ConfigurationSetChannelPeriod(0, 8070))
	parseMessage(port)

	port.Write(messages.ControlOpenChannel(0))

	// port.Write(ant.MessageOpenRXScanMode(0))

	for {
		parseMessage(port)

	}

	// alloc = make([]byte, 55)
	// port.Read(alloc)
	// ant.DebugPrint(alloc)

}
