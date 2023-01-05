package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type VoltmeterGetStatusRequest struct {
	Id int `json:"id"`
}

type VoltmeterGetStatusResponse struct {
	Id      int     `json:"id"`
	Voltage float64 `json:"voltage"`
}

func (s *ShellyV2) fillVoltmeterMetrics(m *shelly.Metrics) {
	if s.status.VoltmeterStatus == nil {
		return
	}

	for i, voltmeter := range s.status.VoltmeterStatus {
		labels := shelly.LineLabels(s.Shelly, "voltmeter", i)

		m.Voltmeter.WithLabelValues(labels...).Add(voltmeter.Voltage)
	}
}

func (s *ShellyV2) getVoltmeterStatus(status *Status) error {
	for i := 0; true; i++ {
		res := VoltmeterGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Voltmeter.GetStatus",
			Params:         VoltmeterGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.VoltmeterStatus = append(status.VoltmeterStatus, res)
	}
	return nil
}
