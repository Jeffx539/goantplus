package constants

// AntMessageSync defined in 7.1 of antplus documentation
const (
	AntMessageSync = 0xa4

	// Control Messages
	ControlMessageResetSystem    = 0x4A
	ControlMessageOpenChannel    = 0x4B
	ControlMessageCloseChannel   = 0x4C
	ControlMessageRequestMessage = 0x4D

	// Configuration Messages
	ConfigurationMessageAssignChannel           = 0x42
	ConfigurationMessageUnassignChannel         = 0x41
	ConfigurationMessageSetChannelPeriod        = 0x43
	ConfigurationMessageSetChannelSearchTimeout = 0x44
	ConfigurationMessageSetChannelRFFreq        = 0x45
	ConfigurationMessageSetNetworkKey           = 0x46
	ConfigurationMessageSetChannelID            = 0x51
	ConfigurationMessageOpenRxScanMode          = 0x5B
)
