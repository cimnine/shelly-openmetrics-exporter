package shelly

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Power                 *prometheus.GaugeVec
	PowerLimit            *prometheus.GaugeVec
	Current               *prometheus.GaugeVec
	CurrentLimit          *prometheus.GaugeVec
	NeutralCurrent        *prometheus.GaugeVec
	NeutralIxsum          *prometheus.GaugeVec
	NeutralMismatch       *prometheus.GaugeVec
	NeutralValid          *prometheus.GaugeVec
	Voltage               *prometheus.GaugeVec
	VoltageLimit          *prometheus.GaugeVec
	PowerFactor           *prometheus.GaugeVec
	Total                 *prometheus.CounterVec
	TotalReturned         *prometheus.CounterVec
	Frequency             *prometheus.GaugeVec
	RelativeHumidity      *prometheus.GaugeVec
	Temperature           *prometheus.GaugeVec
	TemperatureDevice     *prometheus.GaugeVec
	Voltmeter             *prometheus.GaugeVec
	RelayOpen             *prometheus.GaugeVec
	RelayOverpowered      *prometheus.GaugeVec
	RelayValid            *prometheus.GaugeVec
	InputState            *prometheus.GaugeVec
	InputEventCount       *prometheus.CounterVec
	InputPercent          *prometheus.GaugeVec
	HasUpdate             *prometheus.GaugeVec
	CloudEnabled          *prometheus.GaugeVec
	CloudConnected        *prometheus.GaugeVec
	WifiConnected         *prometheus.GaugeVec
	WifiSignal            *prometheus.GaugeVec
	DevicePowerVoltage    *prometheus.GaugeVec
	DevicePowerPercentage *prometheus.GaugeVec
	LightBrightness       *prometheus.GaugeVec
	LightState            *prometheus.GaugeVec
	CoverPosition         *prometheus.GaugeVec
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
		PowerLimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_power_limit_w",
				Help: "The power limit value in this instant.",
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
		CurrentLimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_current_limit_a",
				Help: "The current limit value in this instant.",
			},
			lineLabels,
		),
		NeutralCurrent: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_neutral_current_a",
				Help: "The neutral current value in this instant.",
			},
			deviceLabels,
		),
		NeutralIxsum: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_neutral_ixsum_a",
				Help: "The neutral ixsum value in this instant.",
			},
			deviceLabels,
		),
		NeutralMismatch: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_neutral_mismatch",
				Help: "1 if the neutral is mismatched, else 0.",
			},
			deviceLabels,
		),
		NeutralValid: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_neutral_valid",
				Help: "1 if the neutral is valid, else 0.",
			},
			deviceLabels,
		),
		Voltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_voltage_v",
				Help: "The voltage value in this instant.",
			},
			lineLabels,
		),
		VoltageLimit: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_voltage_limit_v",
				Help: "The voltage limit value in this instant.",
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
		Frequency: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_grid_frequency",
				Help: "Last measured network frequency.",
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
		RelativeHumidity: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_humidity_percent",
				Help: "The current relative humidity in percent.",
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
		Voltmeter: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_voltmeter_v",
				Help: "The voltage value of the built-in volt meters / ADCs in this instant.",
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
				Help: "1 if the input state is on, else 0.",
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
		InputPercent: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_input_percent",
				Help: "The input percent value in this instant.",
			},
			lineLabels,
		),
		HasUpdate: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_update_state",
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
				Name: "shelly_wifi_signal_db",
				Help: "The signal strength of the wifi in this instant.",
			},
			append(deviceLabels, "ssid"),
		),
		DevicePowerPercentage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_device_power_percent",
				Help: "The percentage of the device power source in this instant.",
			},
			append(deviceLabels, "ssid"),
		),
		DevicePowerVoltage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_device_power_v",
				Help: "The voltage of the device power source in this instant.",
			},
			append(deviceLabels, "ssid"),
		),
		LightBrightness: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_light_brightness_percent",
				Help: "The percentage of the brightness of the light in this instant.",
			},
			append(deviceLabels, "ssid"),
		),
		LightState: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_light_state",
				Help: "1 if the light is on, else 0.",
			},
			append(deviceLabels, "ssid"),
		),
		CoverPosition: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "shelly_cover_position_percent",
				Help: "The position of the cover in percent in this instant.",
			},
			append(deviceLabels, "ssid"),
		),
	}

	reg.MustRegister(
		m.Power,
		m.PowerLimit,
		m.Current,
		m.CurrentLimit,
		m.NeutralCurrent,
		m.NeutralIxsum,
		m.NeutralMismatch,
		m.NeutralValid,
		m.Voltage,
		m.VoltageLimit,
		m.PowerFactor,
		m.Total,
		m.TotalReturned,
		m.RelativeHumidity,
		m.Temperature,
		m.TemperatureDevice,
		m.Frequency,
		m.Voltmeter,
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
		m.DevicePowerPercentage,
		m.DevicePowerVoltage,
		m.LightBrightness,
		m.LightState,
		m.CoverPosition,
	)

	return m
}
