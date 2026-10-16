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
	body, err := os.ReadFile("status.json")
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
	want := &Status{
		ActionsStats: &ActionStats{}, Cloud: &CloudStatus{Connected: true, Enabled: true},
		EmeterN: &EmeterN{Ixsum: 0.67},
		Emeters: []Emeter{
			{Current: 0.72, IsValid: true, PowerFactor: 0.69, Power: 113.63, Total: 191131.6, Voltage: 230.09},
			{Current: 0.26, IsValid: true, PowerFactor: 0.64, Power: 38.24, Total: 103767.7, Voltage: 230.75},
			{Current: 0.27, IsValid: true, PowerFactor: 0.5, Power: 30.59, Total: 128226.6, Voltage: 230.74},
		},
		FsFree: 154365, FsMounted: true, FsSize: 233681, Mac: "111CAB11A11E", Mqtt: &MQTTStatus{},
		RamFree: 31884, RamTotal: 49928, Relays: []Relay{{IsValid: true, Source: "input"}},
		Serial: 2414, Time: "07:32", TotalPower: 182.46, Temperature: 29.35,
		Temperatures:      &Temperature{Celsius: 29.35, Fahrenheit: 84.83, IsValid: true},
		TemperatureStatus: "Normal", Unixtime: 1672813941,
		Update: &UpdateStatus{NewVersion: "20221027-110030/v1.12.1-ga9117d3", OldVersion: "20221027-110030/v1.12.1-ga9117d3", Status: "idle"},
		Uptime: 141090, VData: 1, WifiSta: &WifiStatus{Connected: true, Ip: "192.168.1.1", Rssi: -56, Ssid: "Wifi Name"},
	}
	if !reflect.DeepEqual(client.status, want) {
		t.Errorf("status = %+v, want %+v", client.status, want)
	}
}

func TestSettingsFixture(t *testing.T) {
	body, err := os.ReadFile("settings.json")
	if err != nil {
		t.Fatal(err)
	}
	var got Settings
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatal(err)
	}
	want := Settings{
		Actions:   Actions{Names: []string{"out_on_url", "out_off_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "over_power_url", "under_power_url", "n_mismatch_url"}},
		ApRoaming: APRoaming{Enabled: true, Threshold: -55},
		BuildInfo: BuildInfo{BuildId: "20221027-110030/v1.12.1-ga9117d3", BuildTimestamp: time.Date(2022, 10, 27, 11, 0, 30, 0, time.UTC), BuildVersion: "1.0"},
		CfOutput:  100, Cloud: CloudSettings{Connected: true, Enabled: true}, Coiot: CoIOT{Enabled: true, UpdatePeriod: 15},
		Device:  Device{Hostname: "shellyem3-111CAB11A11E", Mac: "111CAB11A11E", NumEmeters: 3, NumOutputs: 1, ReportPeriod: 1, Type: "SHEM-3"},
		EmeterN: EmeterNSettings{MismatchThreshold: 1, RangeExtender: 1},
		Emeters: []EmeterSettings{
			{ApplianceType: "General", Name: "Channel 1", RangeExtender: 1},
			{ApplianceType: "General", Name: "Channel 2", RangeExtender: 1},
			{ApplianceType: "General", Name: "Channel 3", RangeExtender: 1},
		},
		Firmware: "20221027-110030/v1.12.1-ga9117d3", Hwinfo: HardwareInfo{HwRevision: "dev-prototype"},
		Lat: 0.495209, Lng: 0.71693, Login: LoginSettings{Username: "admin"},
		Mqtt: MQTTSettings{CleanSession: true, Id: "shellyem3-111CAB11A11E", KeepAlive: 60, ReconnectTimeoutMax: 60, ReconnectTimeoutMin: 2, Server: "192.168.33.3:1883", UpdatePeriod: 30},
		Name: "shelly-device-name", Relays: []RelaySettings{{DefaultState: "last", ScheduleRules: []interface{}{}}},
		Sntp: SNTPSettings{Enabled: true, Server: "time.google.com"}, Time: "22:30", Timezone: "Europe/Zurich",
		TzDstAuto: true, TzUtcOffset: 3600, Tzautodetect: true, Unixtime: 1672781440,
		WifiAp:   WifiAPSettings{Ssid: "shellyem3-111CAB11A11E"},
		WifiSta:  WifiSettings{Enabled: true, Ipv4Method: "dhcp", Ssid: "Wifi Name 1"},
		WifiSta1: WifiSettings{Enabled: true, Ipv4Method: "dhcp", Ssid: "Wifi Name 2"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("settings = %+v, want %+v", got, want)
	}
}
