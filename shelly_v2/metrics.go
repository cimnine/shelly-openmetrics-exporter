package shelly_v2

import (
	"strconv"

	"shelly-prometheus-exporter/shelly_metrics"
)

type Status struct {
	Switches []SwitchGetStatusResponse
}

func ParseMetrics(status Status, m *shelly_metrics.Metrics) {
	for i, shellySwitch := range status.Switches {
		labelValue := strconv.Itoa(i)

		m.Current.WithLabelValues(labelValue).Set(shellySwitch.Current)
		m.Power.WithLabelValues(labelValue).Set(shellySwitch.Power)
		m.PowerFactor.WithLabelValues(labelValue).Set(shellySwitch.PowerFactor)
		m.Total.WithLabelValues(labelValue).Add(shellySwitch.ActiveEnergy.Total)
		m.Voltage.WithLabelValues(labelValue).Add(shellySwitch.Voltage)
	}
}
