package shelly_v2

import (
	"github.com/cimnine/shelly-openmetrics-exporter/shelly"
)

type WifiGetStatusResponse struct {
	RSSI   int    `json:"rssi"`
	SSID   string `json:"ssid"`
	IP     string `json:"sta_ip"`
	Status string `json:"status"`
}

func (s *ShellyV2) fillWifiMetrics(m *shelly.Metrics) {
	if s.status.WifiStatus == nil {
		return
	}

	labels := append(shelly.DeviceLabels(s.Shelly), s.status.WifiStatus.SSID)

	m.WifiSignal.WithLabelValues(labels...).Set(float64(s.status.WifiStatus.RSSI))
	m.WifiConnected.WithLabelValues(labels...).Set(shelly.BoolToFloat(s.status.WifiStatus.Status == "got ip"))
}

func (s *ShellyV2) getWifiStatus(status *Status) error {
	res := &WifiGetStatusResponse{}
	request := JsonRpc2Request{
		JsonRpcVersion: "2.0",
		Src:            "shelly-openmetrics-exporter",
		Method:         "Wifi.GetStatus",
	}

	end, err := s.do(request, res)
	if end {
		return nil
	}
	if err != nil {
		return err
	}

	status.WifiStatus = res
	return nil
}
