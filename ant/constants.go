package ant

// AntMessageSync defined in 7.1 of antplus documentation
const (
	AntMessageSync = 0xa4

	//
	AntControlMessageResetSystem = 0x4a
	AntControlMessageRequest     = 0x4d

	AntMessageSetChannelPeriod      = 0x43
	AntMessageSetChannelRFFrequency = 0x45
	AntMessageSetNetworkKey         = 0x46
	AntMessageParamCapability       = 0x54

	// Channel Assignments
	AntMessageAssignChannel   = 0x42
	AntMessageUnassignChannel = 0x41
	AntMessageOpenChannel     = 0x4B
	AntMessageSetChannelID    = 0x51

	AntOpenRXScanMode = 0x5B

	AntNetworkKey = "\xB9\xA5\x21\xFB\xBD\x72\xC3\x45"
)
