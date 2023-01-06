package shelly_v2

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type SwitchGetStatusRequest struct {
	Id int `json:"id"`
}

type EnergyCounter struct {
	Total    float64   `json:"total"`
	ByMinute []float64 `json:"by_minute"`
	MinuteTs int       `json:"minute_ts"`
}

type Temperature struct {
	Celsius    float64 `json:"tC"`
	Fahrenheit float64 `json:"tF"`
}

type SwitchGetStatusResponse struct {
	Id           int           `json:"id"`
	Source       string        `json:"source"`
	Output       bool          `json:"output"`
	Power        float64       `json:"apower"`
	PowerFactor  float64       `json:"pf"`
	Voltage      float64       `json:"voltage"`
	Current      float64       `json:"current"`
	ActiveEnergy EnergyCounter `json:"aenergy"`
	Temperature  *Temperature  `json:"temperature"`
}

type SwitchGetConfigRequest struct {
	Id int `json:"id"`
}

type SwitchGetConfigResponse struct {
	AutoOff      bool    `json:"auto_off"`
	AutoOffDelay float64 `json:"auto_off_delay"`
	AutoOn       bool    `json:"auto_on"`
	AutoOnDelay  float64 `json:"auto_on_delay"`
	Id           int     `json:"id"`
	InMode       string  `json:"in_mode"`
	InitialState string  `json:"initial_state"`
	Name         string  `json:"name"`
	CurrentLimit float64 `json:"current_limit"`
	PowerLimit   float64 `json:"power_limit"`
	VoltageLimit float64 `json:"voltage_limit"`
}

func (s *ShellyV2) fillSwitchMetrics(m *shelly.Metrics) {
	if s.status.SwitchesStatus == nil {
		return
	}

	for i, shellySwitch := range s.status.SwitchesStatus {
		labels := shelly.LineLabels(s.Shelly, "switch", i)

		m.Current.WithLabelValues(labels...).Set(shellySwitch.Current)
		m.Power.WithLabelValues(labels...).Set(shellySwitch.Power)
		m.PowerFactor.WithLabelValues(labels...).Set(shellySwitch.PowerFactor)
		m.Total.WithLabelValues(labels...).Add(shellySwitch.ActiveEnergy.Total)
		m.Voltage.WithLabelValues(labels...).Add(shellySwitch.Voltage)
		m.Temperature.WithLabelValues(labels...).Add(shelly.CelsiusToKelvin(shellySwitch.Temperature.Celsius))
	}
	for i, shellySwitch := range s.status.SwitchesConfig {
		labels := shelly.LineLabels(s.Shelly, "switch", i)

		m.CurrentLimit.WithLabelValues(labels...).Set(shellySwitch.CurrentLimit)
		m.PowerLimit.WithLabelValues(labels...).Set(shellySwitch.PowerLimit)
		m.VoltageLimit.WithLabelValues(labels...).Add(shellySwitch.VoltageLimit)
	}
}

func (s *ShellyV2) getSwitchStatus(status *Status) error {
	for i := 0; true; i++ {
		res := SwitchGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Switch.GetStatus",
			Params:         SwitchGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.SwitchesStatus = append(status.SwitchesStatus, res)
	}
	return nil
}

func (s *ShellyV2) getSwitchConfig(status *Status) error {
	for i := 0; true; i++ {
		res := SwitchGetConfigResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Switch.GetConfig",
			Params:         SwitchGetConfigRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.SwitchesConfig = append(status.SwitchesConfig, res)
	}
	return nil
}
