package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type Status struct {
	Switches []SwitchGetStatusResponse
}

func (s *ShellyV2) FillMetrics(m *shelly.Metrics) {
	if s.status.Switches == nil {
		return
	}

	s.readSwitchStatus(m)
}

func (s *ShellyV2) readSwitchStatus(m *shelly.Metrics) {
	if s.status.Switches == nil {
		return
	}

	for i, shellySwitch := range s.status.Switches {
		labels := shelly.LineLabels(s.Shelly, i)

		m.Current.WithLabelValues(labels...).Set(shellySwitch.Current)
		m.Power.WithLabelValues(labels...).Set(shellySwitch.Power)
		m.PowerFactor.WithLabelValues(labels...).Set(shellySwitch.PowerFactor)
		m.Total.WithLabelValues(labels...).Add(shellySwitch.ActiveEnergy.Total)
		m.Voltage.WithLabelValues(labels...).Add(shellySwitch.Voltage)
		m.Temperature.WithLabelValues(labels...).Add(shellySwitch.Temperature.Celsius + 273.15)
	}
}
