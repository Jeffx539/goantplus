package main

import (
	"fmt"
	"time"

	"github.com/jeffx539/goantplus/pkg"
	"github.com/jeffx539/goantplus/pkg/constants"
	"github.com/jeffx539/goantplus/pkg/profiles"
)

func testCB(pl []byte) {
	fmt.Println("Register CB")
	fmt.Println(pl)

	var hr profiles.HeartRate
	hr.UnMarshal(pl)

	fmt.Println(hr)

}

func main() {
	fmt.Println("ANTPlus Dump")

	dev := pkg.MakeDevice("/dev/ttyUSB0")
	if dev == nil {
		panic("Unable to allocate device")
	}

	fmt.Println("Initialised device.")

	dev.InitialiseDevice()
	go dev.Loop()

	c := dev.IntialiseChannel(0x00)
	c.SetNetworkKey(([]byte)(constants.ANTNetworkKey))
	c.AssignChannel(0x40)
	c.SetChannelPeriod(8070)
	c.SetRFFrequency(0x39)
	c.SetChannelID(0x00, 0x00, 0x00, 0x00)
	c.RegisterCallback(testCB)
	c.Open()

	for {
		time.Sleep(1)
	}
}
