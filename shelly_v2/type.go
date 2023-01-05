package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type ShellyV2 struct {
	*shelly.Shelly
	status *Status
}

func New(targetHost string) *ShellyV2 {
	return &ShellyV2{
		Shelly: &shelly.Shelly{TargetHost: targetHost},
	}
}
