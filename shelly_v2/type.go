package shelly_v2

type ShellyV2 struct {
	targetHost string
	status     *Status
}

func New(targetHost string) *ShellyV2 {
	return &ShellyV2{targetHost: targetHost}
}
