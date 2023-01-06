package shelly_v2

import (
	"gitlab.com/cimnine/shelly-prometheus-exporter/shelly"
)

type CoverGetStatusRequest struct {
	Id int `json:"id"`
}

type CoverGetStatusResponse struct {
	Id              int           `json:"id"`
	Source          string        `json:"source"`
	State           string        `json:"state"`
	ActivePower     float64       `json:"apower"`
	Voltage         float64       `json:"voltage"`
	Current         float64       `json:"current"`
	PowerFactor     float64       `json:"pf"`
	ActiveEnergy    EnergyCounter `json:"aenergy"`
	Temperature     Temperature   `json:"temperature"`
	PosControl      bool          `json:"pos_control"`
	CurrentPosition float64       `json:"current_pos"`
}

type CoverGetConfigRequest struct {
	Id int `json:"id"`
}

type ObstructionDetection struct {
	Enable         bool   `json:"enable"`
	Direction      string `json:"direction"`
	Action         string `json:"action"`
	PowerThreshold int    `json:"power_thr"`
	Holdoff        int    `json:"holdoff"`
}

type SafetySwitch struct {
	Enable      bool        `json:"enable"`
	Direction   string      `json:"direction"`
	Action      string      `json:"action"`
	AllowedMove interface{} `json:"allowed_move"`
}

type Motor struct {
	IdlePowerThreshold       int     `json:"idle_power_thr"`
	IdleConfirmPeriodSeconds float64 `json:"idle_confirm_period"`
}

type CoverGetConfigResponse struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Motor            Motor   `json:"motor"`
	MaxTimeOpen      int     `json:"maxtime_open"`
	MaxTimeClose     int     `json:"maxtime_close"`
	InitialState     string  `json:"initial_state"`
	InvertDirections bool    `json:"invert_directions"`
	InMode           string  `json:"in_mode"`
	SwapInputs       bool    `json:"swap_inputs"`
	PowerLimit       float64 `json:"power_limit"`
	VoltageLimit     float64 `json:"voltage_limit"`
	CurrentLimit     float64 `json:"current_limit"`

	SafetySwitch         SafetySwitch         `json:"safety_switch"`
	ObstructionDetection ObstructionDetection `json:"obstruction_detection"`
}

func (s *ShellyV2) fillCoverMetrics(m *shelly.Metrics) {
	if s.status.CoversStatus == nil {
		return
	}

	for i, cover := range s.status.CoversStatus {
		labels := shelly.LineLabels(s.Shelly, "cover", i)

		m.Current.WithLabelValues(labels...).Set(cover.Current)
		m.Power.WithLabelValues(labels...).Set(cover.ActivePower)
		m.PowerFactor.WithLabelValues(labels...).Set(cover.PowerFactor)
		m.Total.WithLabelValues(labels...).Add(cover.ActiveEnergy.Total)
		m.Voltage.WithLabelValues(labels...).Add(cover.Voltage)
		m.Temperature.WithLabelValues(labels...).Add(shelly.CelsiusToKelvin(cover.Temperature.Celsius))
	}
	for i, cover := range s.status.CoversConfig {
		labels := shelly.LineLabels(s.Shelly, "cover", i)

		m.CurrentLimit.WithLabelValues(labels...).Set(cover.CurrentLimit)
		m.PowerLimit.WithLabelValues(labels...).Set(cover.PowerLimit)
		m.VoltageLimit.WithLabelValues(labels...).Add(cover.VoltageLimit)
	}
}

func (s *ShellyV2) getCoverStatus(status *Status) error {
	for i := 0; true; i++ {
		res := CoverGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Cover.GetStatus",
			Params:         CoverGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.CoversStatus = append(status.CoversStatus, res)
	}
	return nil
}

func (s *ShellyV2) getCoverConfig(status *Status) error {
	for i := 0; true; i++ {
		res := CoverGetConfigResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-prometheus-exporter",
			Method:         "Cover.GetConfig",
			Params:         CoverGetConfigRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.CoversConfig = append(status.CoversConfig, res)
	}
	return nil
}
