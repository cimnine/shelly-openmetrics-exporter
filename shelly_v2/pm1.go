package shelly_v2

import "github.com/cimnine/shelly-openmetrics-exporter/shelly"

type PM1GetStatusRequest struct {
	Id int `json:"id"`
}

type PM1GetStatusResponse struct {
	Id                 int           `json:"id"`
	Voltage            float64       `json:"voltage"`
	Current            float64       `json:"current"`
	Power              float64       `json:"apower"`
	ActiveEnergy       EnergyCounter `json:"aenergy"`
	ActiveEnergyReturn EnergyCounter `json:"ret_aenergy"`
	Freq               float64       `json:"freq"`
}

func (s *ShellyV2) fillPM1Metrics(m *shelly.Metrics) {
	if s.status.PM1Status == nil {
		return
	}

	for i, measurement := range s.status.PM1Status {
		labels := shelly.LineLabels(s.Shelly, "meter", i)

		m.Voltage.WithLabelValues(labels...).Add(measurement.Voltage)
		m.Current.WithLabelValues(labels...).Add(measurement.Current)
		m.Power.WithLabelValues(labels...).Add(measurement.Power)
		m.Total.WithLabelValues(labels...).Add(measurement.ActiveEnergy.Total)
		m.TotalReturned.WithLabelValues(labels...).Add(measurement.ActiveEnergyReturn.Total)
		m.Frequency.WithLabelValues(labels...).Add(measurement.Freq)
	}
}

// getPM1Status retrieves power measurement metrics from the PM1 component.
// See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/PM1 for additional information
func (s *ShellyV2) getPM1Status(status *Status) error {
	for i := 0; true; i++ {
		res := PM1GetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "PM1.GetStatus",
			Params:         PM1GetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.PM1Status = append(status.PM1Status, res)
	}
	return nil
}
