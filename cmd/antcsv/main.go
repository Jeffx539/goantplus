package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jeffx539/goantplus/pkg"
	"github.com/jeffx539/goantplus/pkg/constants"
	"github.com/jeffx539/goantplus/pkg/messages"
)

var pathFlag = flag.String("device", "/dev/ttyUSB0", "Path to the ANT+ device")

func heartrateCallback(pl []byte) {
	messages.DebugPrint(pl)
	// var hr profiles.HeartRate
	// hr.UnMarshal(pl)

	// fmt.Fprintf(os.Stdout, "%d,%d\n", time.Now().UnixNano(), hr.HeartRate)

}

func init() {

	flag.StringVar(pathFlag, "d", "/dev/ttyUSB0", "ANT+ Device path")
}

func main() {

	flag.Parse()
	dev := pkg.MakeDevice(*pathFlag)
	if dev == nil {
		panic("Unable to allocate device")
	}

	fmt.Fprintf(os.Stderr, "Initialised Device %s\n", *pathFlag)

	dev.InitialiseDevice()
	go dev.Loop()

	c1 := dev.IntialiseChannel((byte)(dev.GetNextAvailableChannel()))
	c1.SetNetworkKey(([]byte)(constants.ANTNetworkKey))
	c1.AssignChannel(0x00)
	c1.SetChannelPeriod(8070)
	c1.SetRFFrequency(57)
	c1.SetChannelID(0x00, 40, 0x00)
	c1.RegisterCallback(heartrateCallback)
	c1.Open()

	time.Sleep(2000)
	c2 := dev.IntialiseChannel((byte)(dev.GetNextAvailableChannel()))
	c2.SetNetworkKey(([]byte)(constants.ANTNetworkKey))
	c2.AssignChannel(0x00)
	c2.SetChannelPeriod(8070)
	c2.SetRFFrequency(57)
	c2.SetChannelID(0x00, 120, 0x00)
	c2.RegisterCallback(heartrateCallback)
	c2.Open()

	for {
		time.Sleep(1)
	}
}
