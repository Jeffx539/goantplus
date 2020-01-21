package profiles

import "fmt"

type HeartRate struct {
	pageNumber     byte
	PageSpecific0  byte
	PageSpecific1  byte
	PageSpecific2  byte
	HRLSB          byte
	HRMSB          byte
	heartBeatCount byte
	heartRate      byte
}

func (h *HeartRate) UnMarshal(payload []byte) {
	h.pageNumber = payload[2]
	h.PageSpecific0 = payload[3]
	h.PageSpecific1 = payload[4]
	h.PageSpecific2 = payload[5]
	h.HRLSB = payload[6]
	h.HRMSB = payload[7]
	h.heartBeatCount = payload[8]
	h.heartRate = payload[9]
}

func (h HeartRate) String() string {
	return fmt.Sprintf("[HRM] Heartrate %d ", h.heartRate)
}
