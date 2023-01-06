package shelly_v2

import (
	"net/http"

	"github.com/cimnine/shelly-openmetrics-exporter/shelly"
)

type ShellyV2 struct {
	*shelly.Shelly
	status        *Status
	nextMessageID *int
	client        *http.Client
}

func New(targetHost, userAgent, password string) *ShellyV2 {
	initialMessageId := 0
	return &ShellyV2{
		Shelly: &shelly.Shelly{
			TargetHost: targetHost,
			UserAgent:  userAgent,
			Username:   "admin", // must be admin, per documentation
			Password:   password,
		},
		nextMessageID: &initialMessageId,
		client:        &http.Client{},
	}
}
