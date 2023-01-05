package shelly_v1

type ShellyV1 struct {
	targetHost string
	status     *Status
}

func New(targetHost string) *ShellyV1 {
	return &ShellyV1{targetHost: targetHost}
}
