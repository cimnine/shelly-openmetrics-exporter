package shelly_v1

import (
	"strconv"

	"shelly-prometheus-exporter/shelly_metrics"
)

func ParseMetrics(status Status, m *shelly_metrics.Metrics) {
	for i, emeter := range status.Emeters {
		labelValue := strconv.Itoa(i)
		m.Current.WithLabelValues(labelValue).Set(emeter.Current)
		m.Power.WithLabelValues(labelValue).Set(emeter.Power)
		m.PowerFactor.WithLabelValues(labelValue).Set(emeter.PowerFactor)
		m.Total.WithLabelValues(labelValue).Add(emeter.Total)
		m.TotalReturned.WithLabelValues(labelValue).Add(emeter.TotalReturned)
		m.Voltage.WithLabelValues(labelValue).Add(emeter.Voltage)
	}
}
