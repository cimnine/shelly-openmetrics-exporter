package shelly_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Power         *prometheus.GaugeVec
	Current       *prometheus.GaugeVec
	Voltage       *prometheus.GaugeVec
	PowerFactor   *prometheus.GaugeVec
	Total         *prometheus.CounterVec
	TotalReturned *prometheus.CounterVec
}

func NewMetrics(reg *prometheus.Registry) *Metrics {
	commonLabels := []string{"line"}
	m := &Metrics{
		Power: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_power_w",
				Help: "The power value in this instant.",
			},
			commonLabels,
		),
		Current: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_current_a",
				Help: "The current value in this instant.",
			},
			commonLabels,
		),
		Voltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_voltage_v",
				Help: "The voltage value in this instant.",
			},
			commonLabels,
		),
		PowerFactor: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_pf",
				Help: "The power factor in this instant.",
			},
			commonLabels,
		),
		Total: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "shelly_total_wh",
				Help: "The total consumed energy up to this instant.",
			},
			commonLabels,
		),
		TotalReturned: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "shelly_total_returned_wh",
				Help: "The total returned energy up to this instant.",
			},
			commonLabels,
		),
	}

	reg.MustRegister(
		m.Power,
		m.Current,
		m.Voltage,
		m.PowerFactor,
		m.Total,
		m.TotalReturned,
	)

	return m
}
