package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type Status struct {
	SwitchesStatus    []SwitchGetStatusResponse
	SwitchesConfig    []SwitchGetConfigResponse
	InputStatus       []InputGetStatusResponse
	WifiStatus        *WifiGetStatusResponse
	CloudStatus       *CloudGetStatusResponse
	CloudConfig       *CloudGetConfigResponse
	VoltmeterStatus   []VoltmeterGetStatusResponse
	TemperatureStatus []TemperatureGetStatusResponse
	HumidityStatus    []HumidityGetStatusResponse
	DevicePowerStatus []DevicePowerGetStatusResponse
	LightStatus       []LightGetStatusResponse
	CoversConfig      []CoverGetConfigResponse
	CoversStatus      []CoverGetStatusResponse
}

func (s *ShellyV2) FillMetrics(m *shelly.Metrics) {
	s.fillSwitchMetrics(m)
	s.fillInputMetrics(m)
	s.fillWifiMetrics(m)
	s.fillCloudMetrics(m)
	s.fillVoltmeterMetrics(m)
	s.fillHumidityMetrics(m)
	s.fillTemperatureMetrics(m)
	s.fillDevicePowerMetrics(m)
	s.fillCoverMetrics(m)
}
