package shelly_v2

import (
	"net/http"

	"shelly-prometheus-exporter/shelly"
)

type ShellyV2 struct {
	*shelly.Shelly
	status        *Status
	nextMessageID *int
	client        *http.Client
}

func New(targetHost string) *ShellyV2 {
	initialMessageId := 0
	return &ShellyV2{
		Shelly:        &shelly.Shelly{TargetHost: targetHost},
		nextMessageID: &initialMessageId,
		client:        &http.Client{},
	}
}
