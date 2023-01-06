package shelly_v1

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type ShellyV1 struct {
	*shelly.Shelly
	status *Status
}

func New(targetHost, userAgent, username, password string) *ShellyV1 {
	return &ShellyV1{
		Shelly: &shelly.Shelly{
			TargetHost: targetHost,
			UserAgent:  userAgent,
			Username:   username,
			Password:   password,
		},
	}
}
