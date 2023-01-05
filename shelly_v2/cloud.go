package shelly_v2

import (
	"shelly-prometheus-exporter/shelly"
)

type CloudGetStatusResponse struct {
	Connected bool `json:"connected"`
}

type CloudGetConfigResponse struct {
	Enable bool   `json:"enable"`
	Server string `json:"server"`
}

func (s *ShellyV2) fillCloudMetrics(m *shelly.Metrics) {
	if s.status.CloudStatus == nil {
		return
	}

	labels := shelly.DeviceLabels(s.Shelly)

	m.CloudConnected.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.CloudStatus.Connected))
	m.CloudEnabled.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.CloudConfig.Enable))
}

func (s *ShellyV2) getCloudConfig(status *Status) error {
	res := &CloudGetConfigResponse{}
	request := JsonRpc2Request{
		JsonRpcVersion: "2.0",
		Src:            "shelly-prometheus-exporter",
		Method:         "Cloud.GetConfig",
	}

	end, err := s.do(request, res)
	if end {
		return nil
	}
	if err != nil {
		return err
	}

	status.CloudConfig = res
	return nil
}

func (s *ShellyV2) getCloudStatus(status *Status) error {
	res := &CloudGetStatusResponse{}
	request := JsonRpc2Request{
		JsonRpcVersion: "2.0",
		Src:            "shelly-prometheus-exporter",
		Method:         "Cloud.GetStatus",
	}

	end, err := s.do(request, res)
	if end {
		return nil
	}
	if err != nil {
		return err
	}

	status.CloudStatus = res
	return nil
}
