package shelly_v1

import (
	"shelly-prometheus-exporter/shelly"
)

func (s *ShellyV1) FillMetrics(m *shelly.Metrics) {
	s.readEmeters(m)
	s.readTemperature(m)
	s.readAdcs(m)
	s.readRelays(m)
	s.readInputs(m)
	s.readUpdate(m)
	s.readCloud(m)
	s.readWifi(m)
}

func (s *ShellyV1) readWifi(m *shelly.Metrics) {
	if s.status.WifiSta == nil {
		return
	}

	labels := append(shelly.DeviceLabels(s.Shelly), s.status.WifiSta.Ssid)
	m.WifiConnected.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.WifiSta.Connected))
	m.WifiSignal.WithLabelValues(labels...).Set(float64(s.status.WifiSta.Rssi))
}

func (s *ShellyV1) readCloud(m *shelly.Metrics) {
	if s.status.Cloud == nil {
		return
	}

	labels := shelly.DeviceLabels(s.Shelly)
	m.CloudEnabled.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.Cloud.Enabled))
	m.CloudConnected.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.Cloud.Connected))
}

func (s *ShellyV1) readUpdate(m *shelly.Metrics) {
	if s.status.Update == nil {
		return
	}

	labels := shelly.DeviceLabels(s.Shelly)
	m.HasUpdate.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.HasUpdate))
}

func (s *ShellyV1) readInputs(m *shelly.Metrics) {
	if s.status.Relays == nil {
		return
	}

	for i, input := range s.status.Inputs {
		labels := shelly.LineLabels(s.Shelly, i)
		m.InputState.WithLabelValues(labels...).Set(float64(input.Input))
		m.InputEventCount.WithLabelValues(labels...).Add(float64(input.EventCount))
	}
}

func (s *ShellyV1) readRelays(m *shelly.Metrics) {
	if s.status.Relays == nil {
		return
	}

	for i, relay := range s.status.Relays {

		labels := shelly.LineLabels(s.Shelly, i)
		m.RelayOpen.WithLabelValues(labels...).Set(shelly.BoolToFloat(relay.IsOn))
		m.RelayOverpowered.WithLabelValues(labels...).Set(shelly.BoolToFloat(relay.IsOn))
		m.RelayValid.WithLabelValues(labels...).Set(shelly.BoolToFloat(relay.IsValid))
	}
}

func (s *ShellyV1) readAdcs(m *shelly.Metrics) {
	if s.status.ADCs == nil {
		return
	}

	for i, adc := range s.status.ADCs {
		labels := shelly.LineLabels(s.Shelly, i)
		m.Current.WithLabelValues(labels...).Set(adc.Voltage)
	}
}

func (s *ShellyV1) readTemperature(m *shelly.Metrics) {
	if s.status.Temperatures == nil {
		return
	}

	labels := shelly.DeviceLabels(s.Shelly)
	m.TemperatureDevice.WithLabelValues(labels...).Add(s.status.Temperatures.Celsius)
}

func (s *ShellyV1) readEmeters(m *shelly.Metrics) {
	if s.status.Emeters == nil {
		return
	}

	for i, emeter := range s.status.Emeters {
		labels := shelly.LineLabels(s.Shelly, i)

		m.Current.WithLabelValues(labels...).Set(emeter.Current)
		m.Power.WithLabelValues(labels...).Set(emeter.Power)
		m.PowerFactor.WithLabelValues(labels...).Set(emeter.PowerFactor)
		m.Total.WithLabelValues(labels...).Add(emeter.Total)
		m.TotalReturned.WithLabelValues(labels...).Add(emeter.TotalReturned)
		m.Voltage.WithLabelValues(labels...).Add(emeter.Voltage)
	}
}
