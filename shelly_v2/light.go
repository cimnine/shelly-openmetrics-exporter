package shelly_v2

import (
	"github.com/cimnine/shelly-openmetrics-exporter/shelly"
)

type LightGetStatusRequest struct {
	Id int `json:"id"`
}

type LightGetStatusResponse struct {
	Id             int     `json:"id"`
	Source         string  `json:"source"`
	Output         bool    `json:"output"`
	Brightness     float64 `json:"brightness"`
	TimerStartedAt float64 `json:"timer_started_at"`
	TimerDuration  int64   `json:"timer_duration"`
}

func (s *ShellyV2) fillLightMetrics(m *shelly.Metrics) {
	if s.status.LightStatus == nil {
		return
	}

	for i, light := range s.status.LightStatus {
		labels := shelly.LineLabels(s.Shelly, "light", i)

		m.LightBrightness.WithLabelValues(labels...).Add(light.Brightness)
		m.LightState.WithLabelValues(labels...).Add(shelly.BoolToFloat(light.Output))
	}
}

func (s *ShellyV2) getLightStatus(status *Status) error {
	for i := 0; true; i++ {
		res := LightGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "Light.GetStatus",
			Params:         LightGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.LightStatus = append(status.LightStatus, res)
	}
	return nil
}
