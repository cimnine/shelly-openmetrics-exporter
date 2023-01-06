package shelly_v2

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type TemperatureGetStatusRequest struct {
	Id int `json:"id"`
}

type TemperatureGetStatusResponse struct {
	Id         int     `json:"id"`
	Celsius    float64 `json:"tC"`
	Fahrenheit float64 `json:"tF"`
}

func (s *ShellyV2) fillTemperatureMetrics(m *shelly.Metrics) {
	if s.status.TemperatureStatus == nil {
		return
	}

	for i, temperature := range s.status.TemperatureStatus {
		labels := shelly.LineLabels(s.Shelly, "temperature", i)

		m.Temperature.WithLabelValues(labels...).Add(shelly.CelsiusToKelvin(temperature.Celsius))
	}
}

func (s *ShellyV2) getTemperatureStatus(status *Status) error {
	for i := 0; true; i++ {
		res := TemperatureGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Temperature.GetStatus",
			Params:         TemperatureGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.TemperatureStatus = append(status.TemperatureStatus, res)
	}
	return nil
}
