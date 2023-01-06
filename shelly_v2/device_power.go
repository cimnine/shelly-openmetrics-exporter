package shelly_v2

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type DevicePowerGetStatusRequest struct {
	Id int `json:"id"`
}

type Battery struct {
	Voltage float64 `json:"V"`
	Percent float64 `json:"percent"`
}

type DevicePowerGetStatusResponse struct {
	Id       int      `json:"id"`
	Battery  *Battery `json:"battery"`
	External struct {
		Present bool `json:"present"`
	} `json:"external"`
}

func (s *ShellyV2) fillDevicePowerMetrics(m *shelly.Metrics) {
	if s.status.DevicePowerStatus == nil {
		return
	}

	for i, devicePower := range s.status.DevicePowerStatus {
		labels := shelly.LineLabels(s.Shelly, "devicePower", i)

		if devicePower.Battery == nil {
			continue
		}
		m.DevicePowerVoltage.WithLabelValues(labels...).Add(devicePower.Battery.Voltage)
		m.DevicePowerPercentage.WithLabelValues(labels...).Add(devicePower.Battery.Percent)
	}
}

func (s *ShellyV2) getDevicePowerStatus(status *Status) error {
	for i := 0; true; i++ {
		res := DevicePowerGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "DevicePower.GetStatus",
			Params:         DevicePowerGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.DevicePowerStatus = append(status.DevicePowerStatus, res)
	}
	return nil
}
