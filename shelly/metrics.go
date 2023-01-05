package shelly

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
	ADCVoltage        *prometheus.GaugeVec
	RelayOpen         *prometheus.GaugeVec
	RelayOverpowered  *prometheus.GaugeVec
	RelayValid        *prometheus.GaugeVec
	InputState        *prometheus.GaugeVec
	InputEventCount   *prometheus.CounterVec
	HasUpdate         *prometheus.GaugeVec
	CloudEnabled      *prometheus.GaugeVec
	CloudConnected    *prometheus.GaugeVec
	WifiConnected     *prometheus.GaugeVec
	WifiSignal        *prometheus.GaugeVec
}

func NewMetrics(reg *prometheus.Registry) *Metrics {
	deviceLabels := []string{"target"}
	lineLabels := append(deviceLabels, "line")

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
			deviceLabels,
		),
		ADCVoltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_adc_voltage_v",
				Help: "The voltage value of the built-in ADCs in this instant.",
			},
			lineLabels,
		),
		RelayOpen: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_relay_open",
				Help: "1 if the relay is open, else 0.",
			},
			lineLabels,
		),
		RelayOverpowered: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_relay_overpowered",
				Help: "1 if the relay is overpowered, else 0.",
			},
			lineLabels,
		),
		RelayValid: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_relay_valid",
				Help: "1 if the relay is valid, else 0.",
			},
			lineLabels,
		),
		InputState: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_input_state",
				Help: "The input value in this instant.",
			},
			lineLabels,
		),
		InputEventCount: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "shelly_input_event_count",
				Help: "The current value of the event_cnt counter.",
			},
			lineLabels,
		),
		HasUpdate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_has_update",
				Help: "1 if the shelly can be updated, else 0.",
			},
			deviceLabels,
		),
		CloudEnabled: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_cloud_enabled",
				Help: "1 if the shelly cloud is enabled, else 0.",
			},
			deviceLabels,
		),
		CloudConnected: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_cloud_connected",
				Help: "1 if the shelly is connected to the cloud, else 0.",
			},
			deviceLabels,
		),
		WifiConnected: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_wifi_connected",
				Help: "1 if the shelly is connected to a wifi, else 0.",
			},
			append(deviceLabels, "ssid"),
		),
		WifiSignal: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_wifi_signal",
				Help: "the signal strength of the wifi",
			},
			append(deviceLabels, "ssid"),
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
		m.ADCVoltage,
		m.RelayOpen,
		m.RelayOverpowered,
		m.RelayValid,
		m.InputState,
		m.InputEventCount,
		m.HasUpdate,
		m.CloudEnabled,
		m.CloudConnected,
		m.WifiConnected,
		m.WifiSignal,
	)

	return m
}
