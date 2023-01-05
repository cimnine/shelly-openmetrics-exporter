package shelly_v2

import (
	"strconv"

	"shelly-prometheus-exporter/shelly_metrics"
)

type Status struct {
	Switches []SwitchGetStatusResponse
}

func (s *ShellyV2) FillMetrics(m *shelly_metrics.Metrics) {
	if s.status.Switches == nil {
		return
	}

	for i, shellySwitch := range s.status.Switches {
		line := strconv.Itoa(i)
		s.readSwitchStatus(m, line, shellySwitch)
	}
}

func (s *ShellyV2) readSwitchStatus(m *shelly_metrics.Metrics, line string, shellySwitch SwitchGetStatusResponse) {
	m.Current.WithLabelValues(s.targetHost, line).Set(shellySwitch.Current)
	m.Power.WithLabelValues(s.targetHost, line).Set(shellySwitch.Power)
	m.PowerFactor.WithLabelValues(s.targetHost, line).Set(shellySwitch.PowerFactor)
	m.Total.WithLabelValues(s.targetHost, line).Add(shellySwitch.ActiveEnergy.Total)
	m.Voltage.WithLabelValues(s.targetHost, line).Add(shellySwitch.Voltage)
	m.Temperature.WithLabelValues(s.targetHost, line).Add(shellySwitch.Temperature.Celsius + 273.15)
}
