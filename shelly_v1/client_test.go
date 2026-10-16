package shelly_v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestFetchStatusFixture(t *testing.T) {
	tests := []struct {
		file string
		want Status
	}{
		{"status.json", Status{
			ActionsStats: &ActionStats{}, Cloud: &CloudStatus{Connected: true, Enabled: true},
			CfgChangedCnt: 1, EmeterN: &EmeterN{Ixsum: 0.76},
			Emeters: []Emeter{
				{Current: 0.96, IsValid: true, PowerFactor: 0.66, Power: 146.78, Total: 383339.6, Voltage: 231.60},
				{Current: 0.46, IsValid: true, PowerFactor: 0.70, Power: 74.46, Total: 441525.7, Voltage: 231.41},
				{Current: 0.35, IsValid: true, PowerFactor: 0.51, Power: 41.10, Total: 455871.9, Voltage: 232.50},
			},
			FsFree: 154114, FsMounted: true, FsSize: 233681, Mac: "020000000001", Mqtt: &MQTTStatus{},
			RamFree: 31088, RamTotal: 49920, Relays: []Relay{{IsOn: true, IsValid: true, Source: "input"}},
			Serial: 21697, Time: "20:12", TotalPower: 262.34, Unixtime: 1792174344,
			Update: &UpdateStatus{NewVersion: "20230913-114244/v1.14.0-gcb84623", OldVersion: "20230913-114244/v1.14.0-gcb84623", Status: "idle"},
			Uptime: 19000531, VData: 1, WifiSta: &WifiStatus{Connected: true, Ip: "192.0.2.155", Rssi: -58, Ssid: "Example WiFi"},
		}},
		{"status_second.json", Status{
			ActionsStats: &ActionStats{}, Cloud: &CloudStatus{Connected: true, Enabled: true},
			EmeterN: &EmeterN{Ixsum: 1.94},
			Emeters: []Emeter{
				{Current: 1.02, IsValid: true, PowerFactor: 0.52, Power: 121.15, Total: 6125853.8, Voltage: 231.92},
				{Current: 2.27, IsValid: true, PowerFactor: 0.64, Power: 330.98, Total: 12367543.2, TotalReturned: 219567.4, Voltage: 230.93},
				{Current: 0.22, IsValid: true, PowerFactor: 0.95, Power: 47.93, Total: 699745.3, TotalReturned: 3.6, Voltage: 233.16},
			},
			FsFree: 154114, FsMounted: true, FsSize: 233681, Mac: "020000000002", Mqtt: &MQTTStatus{},
			RamFree: 31152, RamTotal: 49920, Relays: []Relay{{IsValid: true, Source: "input"}},
			Serial: 12156, Time: "20:12", TotalPower: 500.06, Unixtime: 1792174324,
			Update: &UpdateStatus{NewVersion: "20230913-114244/v1.14.0-gcb84623", OldVersion: "20230913-114244/v1.14.0-gcb84623", Status: "idle"},
			Uptime: 13566281, VData: 1, WifiSta: &WifiStatus{Connected: true, Ip: "192.0.2.154", Rssi: -58, Ssid: "Example WiFi"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/status" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(body)
			}))
			defer server.Close()
			client := New(strings.TrimPrefix(server.URL, "http://"), "test", "", "")
			if err := client.FetchStatus(); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(client.status, &tt.want) {
				t.Errorf("status = %+v, want %+v", client.status, tt.want)
			}
		})
	}
}

func TestSettingsFixture(t *testing.T) {
	tests := []struct {
		file         string
		mac          string
		name         string
		threshold    int
		relayOn      bool
		defaultState string
	}{
		{"settings.json", "020000000001", "Power meter 1", -55, true, "last"},
		{"settings_second.json", "020000000002", "Power meter 2", -60, false, "off"},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			var got Settings
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatal(err)
			}
			want := Settings{
				Actions:   Actions{Names: []string{"out_on_url", "out_off_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "n_mismatch_url"}},
				ApRoaming: APRoaming{Enabled: true, Threshold: tt.threshold},
				BuildInfo: BuildInfo{BuildId: "20230913-114244/v1.14.0-gcb84623", BuildTimestamp: time.Date(2023, 9, 13, 11, 42, 44, 0, time.UTC), BuildVersion: "1.0"},
				CfOutput:  100, Cloud: CloudSettings{Connected: true, Enabled: true}, Coiot: CoIOT{Enabled: true, UpdatePeriod: 15},
				Device:  Device{Hostname: "shellyem3-" + tt.mac, Mac: tt.mac, NumEmeters: 3, NumOutputs: 1, ReportPeriod: 1, Type: "SHEM-3"},
				EmeterN: EmeterNSettings{MismatchThreshold: 1, RangeExtender: 1},
				Emeters: []EmeterSettings{
					{ApplianceType: "General", Name: "Channel 1", RangeExtender: 1},
					{ApplianceType: "General", Name: "Channel 2", RangeExtender: 1},
					{ApplianceType: "General", Name: "Channel 3", RangeExtender: 1},
				},
				Firmware: "20230913-114244/v1.14.0-gcb84623", Hwinfo: HardwareInfo{HwRevision: "dev-prototype"},
				Lat: 1, Lng: 2, Login: LoginSettings{Username: "admin"},
				Mqtt: MQTTSettings{CleanSession: true, Id: "shellyem3-" + tt.mac, KeepAlive: 60, ReconnectTimeoutMax: 60, ReconnectTimeoutMin: 2, Server: "192.0.2.1:1883", UpdatePeriod: 30},
				Name: tt.name, Relays: []RelaySettings{{Ison: tt.relayOn, DefaultState: tt.defaultState, ScheduleRules: []interface{}{}}},
				Sntp: SNTPSettings{Enabled: true, Server: "time.example.org"}, Time: "20:12", Timezone: "Etc/UTC",
				TzDstAuto: true, Tzautodetect: true, Unixtime: 1792174343,
				WifiAp:   WifiAPSettings{Ssid: "shellyem3-" + tt.mac},
				WifiSta:  WifiSettings{Enabled: true, Ipv4Method: "dhcp", Ssid: "Example WiFi"},
				WifiSta1: WifiSettings{Enabled: true, Ipv4Method: "dhcp", Ssid: "Example WiFi 2"},
			}
			if !reflect.DeepEqual(got, want) {
				t.Errorf("settings = %+v, want %+v", got, want)
			}
		})
	}
}
