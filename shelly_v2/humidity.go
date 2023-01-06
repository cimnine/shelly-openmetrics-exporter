package shelly_v2

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type HumidityGetStatusRequest struct {
	Id int `json:"id"`
}

type HumidityGetStatusResponse struct {
	Id               int     `json:"id"`
	RelativeHumidity float64 `json:"rh"`
}

func (s *ShellyV2) fillHumidityMetrics(m *shelly.Metrics) {
	if s.status.HumidityStatus == nil {
		return
	}

	for i, humidity := range s.status.HumidityStatus {
		labels := shelly.LineLabels(s.Shelly, "humidity", i)

		m.RelativeHumidity.WithLabelValues(labels...).Add(humidity.RelativeHumidity)
	}
}

func (s *ShellyV2) getHumidityStatus(status *Status) error {
	for i := 0; true; i++ {
		res := HumidityGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Humidity.GetStatus",
			Params:         HumidityGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.HumidityStatus = append(status.HumidityStatus, res)
	}
	return nil
}
