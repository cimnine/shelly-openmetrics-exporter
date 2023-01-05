package shelly_v1

import (
	"strconv"

	"shelly-prometheus-exporter/shelly_metrics"
)

func (s *ShellyV1) FillMetrics(m *shelly_metrics.Metrics) {
	s.readEmeters(m)
	s.readTemperature(m)
}

func (s *ShellyV1) readTemperature(m *shelly_metrics.Metrics) {
	if s.status.Temperatures == nil {
		return
	}

	m.TemperatureDevice.WithLabelValues(s.targetHost).Add(s.status.Temperatures.Celsius)
}

func (s *ShellyV1) readEmeters(m *shelly_metrics.Metrics) {
	if s.status.Emeters == nil {
		return
	}

	for i, emeter := range s.status.Emeters {
		line := strconv.Itoa(i)

		m.Current.WithLabelValues(s.targetHost, line).Set(emeter.Current)
		m.Power.WithLabelValues(s.targetHost, line).Set(emeter.Power)
		m.PowerFactor.WithLabelValues(s.targetHost, line).Set(emeter.PowerFactor)
		m.Total.WithLabelValues(s.targetHost, line).Add(emeter.Total)
		m.TotalReturned.WithLabelValues(s.targetHost, line).Add(emeter.TotalReturned)
		m.Voltage.WithLabelValues(s.targetHost, line).Add(emeter.Voltage)
	}
}
