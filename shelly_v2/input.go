package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type InputGetStatusRequest struct {
	Id int `json:"id"`
}

type InputGetStatusResponse struct {
	Id      int     `json:"id"`
	State   bool    `json:"state"`
	Percent float64 `json:"percent"`
}

func (s *ShellyV2) fillInputMetrics(m *shelly.Metrics) {
	if s.status.InputStatus == nil {
		return
	}

	for i, input := range s.status.InputStatus {
		labels := shelly.LineLabels(s.Shelly, "input", i)

		m.InputState.WithLabelValues(labels...).Set(shelly.BoolToFloat(input.State))
		m.InputPercent.WithLabelValues(labels...).Set(input.Percent)
	}
}

func (s *ShellyV2) getInputStatus(status *Status) error {
	for i := 0; true; i++ {
		res := InputGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Input.GetStatus",
			Params:         InputGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.InputStatus = append(status.InputStatus, res)
	}
	return nil
}
