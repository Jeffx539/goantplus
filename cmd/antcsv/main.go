package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jeffx539/goantplus/pkg"
	"github.com/jeffx539/goantplus/pkg/constants"
	"github.com/jeffx539/goantplus/pkg/profiles"
)

var pathFlag = flag.String("device", "/dev/ttyUSB0", "Path to the ANT+ device")

func heartrateCallback(pl []byte) {
	var hr profiles.HeartRate
	hr.UnMarshal(pl)

	fmt.Fprintf(os.Stdout, "%d,%d\n", time.Now().UnixNano(), hr.HeartRate)

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
	c1.AssignChannel(0x40)
	c1.SetChannelPeriod(8070)
	c1.SetRFFrequency(0x39)
	c1.SetChannelID(0x00, 0x00, 0x00, 0x00)
	c1.RegisterCallback(heartrateCallback)
	c1.Open()

	for {
		time.Sleep(1)
	}
}
