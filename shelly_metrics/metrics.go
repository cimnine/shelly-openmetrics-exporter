package shelly_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Power             *prometheus.GaugeVec
	Current           *prometheus.GaugeVec
	Voltage           *prometheus.GaugeVec
	PowerFactor       *prometheus.GaugeVec
	Total             *prometheus.CounterVec
	TotalReturned     *prometheus.CounterVec
	Temperature       *prometheus.GaugeVec
	TemperatureDevice *prometheus.GaugeVec
}

func NewMetrics(reg *prometheus.Registry) *Metrics {
	lineLabels := []string{"target", "line"}
	m := &Metrics{
		Power: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_power_w",
				Help: "The power value in this instant.",
			},
			lineLabels,
		),
		Current: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_current_a",
				Help: "The current value in this instant.",
			},
			lineLabels,
		),
		Voltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_voltage_v",
				Help: "The voltage value in this instant.",
			},
			lineLabels,
		),
		PowerFactor: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_pf",
				Help: "The power factor in this instant.",
			},
			lineLabels,
		),
		Total: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "shelly_total_wh",
				Help: "The total consumed energy up to this instant.",
			},
			lineLabels,
		),
		TotalReturned: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "shelly_total_returned_wh",
				Help: "The total returned energy up to this instant.",
			},
			lineLabels,
		),
		Temperature: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_temperature_k",
				Help: "The current temperature in degrees of kelvin.",
			},
			lineLabels,
		),
		TemperatureDevice: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_temperature_device_k",
				Help: "The current temperature on the device in degrees of kelvin.",
			},
			[]string{"target"},
		),
	}

	reg.MustRegister(
		m.Power,
		m.Current,
		m.Voltage,
		m.PowerFactor,
		m.Total,
		m.TotalReturned,
		m.Temperature,
		m.TemperatureDevice,
	)

	return m
}
