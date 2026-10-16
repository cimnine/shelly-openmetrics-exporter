package shelly_v2

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestComponentFixtures(t *testing.T) {
	tests := []struct {
		method  string
		fetch   func(*ShellyV2, *Status) error
		want    Status
		indexed bool
	}{
		{"Switch.GetStatus", (*ShellyV2).getSwitchStatus, Status{SwitchesStatus: []SwitchGetStatusResponse{{
			Source: "init", Output: true, Current: 0.163, Voltage: 230.1,
			ActiveEnergy: EnergyCounter{Total: 204.452, ByMinute: []float64{0, 0, 0}, MinuteTs: 1672814023},
			Temperature:  &Temperature{Celsius: 42.6, Fahrenheit: 108.7},
		}}}, true},
		{"Switch.GetConfig", (*ShellyV2).getSwitchConfig, Status{SwitchesConfig: []SwitchGetConfigResponse{{
			AutoOffDelay: 60, AutoOnDelay: 60, InMode: "detached", InitialState: "on", Name: "Channel 1",
			CurrentLimit: 10, PowerLimit: 700, VoltageLimit: 280,
		}}}, true},
		{"Input.GetStatus", (*ShellyV2).getInputStatus, Status{InputStatus: []InputGetStatusResponse{{}}}, true},
		{"Wifi.GetStatus", (*ShellyV2).getWifiStatus, Status{WifiStatus: &WifiGetStatusResponse{
			RSSI: -56, SSID: "Fascht e Familie", IP: "10.10.11.171", Status: "got ip",
		}}, false},
		{"Cloud.GetStatus", (*ShellyV2).getCloudStatus, Status{CloudStatus: &CloudGetStatusResponse{Connected: true}}, false},
		{"Cloud.GetConfig", (*ShellyV2).getCloudConfig, Status{CloudConfig: &CloudGetConfigResponse{
			Enable: true, Server: "shelly-0-eu.shelly.cloud:6022/jrpc",
		}}, false},
		{"Voltmeter.GetStatus", (*ShellyV2).getVoltmeterStatus, Status{VoltmeterStatus: []VoltmeterGetStatusResponse{{Voltage: 3.5}}}, true},
		{"Temperature.GetStatus", (*ShellyV2).getTemperatureStatus, Status{TemperatureStatus: []TemperatureGetStatusResponse{{Celsius: 24.4, Fahrenheit: 75.9}}}, true},
		{"Humidity.GetStatus", (*ShellyV2).getHumidityStatus, Status{HumidityStatus: []HumidityGetStatusResponse{{RelativeHumidity: 73.7}}}, true},
		{"DevicePower.GetStatus", (*ShellyV2).getDevicePowerStatus, Status{DevicePowerStatus: []DevicePowerGetStatusResponse{{Battery: &Battery{Voltage: 4.87, Percent: 50}}}}, true},
		{"Light.GetStatus", (*ShellyV2).getLightStatus, Status{LightStatus: []LightGetStatusResponse{{Source: "timer", Brightness: 50, TimerStartedAt: 1626942399.36, TimerDuration: 60}}}, true},
		{"Cover.GetStatus", (*ShellyV2).getCoverStatus, Status{CoversStatus: []CoverGetStatusResponse{{
			Source: "limit_switch", State: "open", Voltage: 233,
			ActiveEnergy: EnergyCounter{Total: 48.996, ByMinute: []float64{0, 0, 0}, MinuteTs: 1654604045},
			Temperature:  Temperature{Celsius: 55.4, Fahrenheit: 131.7}, PosControl: true, CurrentPosition: 100,
		}}}, true},
		{"Cover.GetConfig", (*ShellyV2).getCoverConfig, Status{CoversConfig: []CoverGetConfigResponse{{
			Name: "myCover", Motor: Motor{IdlePowerThreshold: 2, IdleConfirmPeriodSeconds: 0.25},
			MaxTimeOpen: 60, MaxTimeClose: 60, InitialState: "stopped", InMode: "dual",
			PowerLimit: 2800, VoltageLimit: 280, CurrentLimit: 10,
			SafetySwitch:         SafetySwitch{Direction: "both", Action: "stop"},
			ObstructionDetection: ObstructionDetection{Direction: "open", Action: "stop", PowerThreshold: 120, Holdoff: 1},
		}}}, true},
		{"PM1.GetStatus", (*ShellyV2).getPM1Status, Status{PM1Status: []PM1GetStatusResponse{{
			Voltage: 225.9, Freq: 50,
			ActiveEnergy:       EnergyCounter{Total: 11.679, ByMinute: []float64{0, 0, 0}, MinuteTs: 1654511972},
			ActiveEnergyReturn: EnergyCounter{Total: 4.126, ByMinute: []float64{0, 0, 0}, MinuteTs: 1654511318},
		}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			fixture, err := os.ReadFile(tt.method + ".json")
			if err != nil {
				t.Fatal(err)
			}
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/rpc" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				var request struct {
					ID     int    `json:"id"`
					Method string `json:"method"`
					Params *struct {
						ID int `json:"id"`
					} `json:"params"`
				}
				if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
					t.Error(err)
					http.Error(w, "invalid request", http.StatusBadRequest)
					return
				}
				if request.Method != tt.method {
					t.Errorf("method = %q, want %q", request.Method, tt.method)
				}
				if request.ID != calls {
					t.Errorf("message ID = %d, want %d", request.ID, calls)
				}
				if tt.indexed {
					if request.Params == nil || request.Params.ID != calls {
						t.Errorf("component params = %+v, want id %d", request.Params, calls)
					}
				} else if request.Params != nil {
					t.Errorf("unexpected params: %+v", request.Params)
				}
				calls++
				w.Header().Set("Content-Type", "application/json")
				if calls > 1 {
					_, _ = w.Write([]byte(`{"error":{"code":-105,"message":"No such component"}}`))
					return
				}
				if err := json.NewEncoder(w).Encode(struct {
					ID     int             `json:"id"`
					Result json.RawMessage `json:"result"`
				}{request.ID, fixture}); err != nil {
					t.Error(err)
				}
			}))
			defer server.Close()
			client := New(strings.TrimPrefix(server.URL, "http://"), "test", "")
			var got Status
			if err := tt.fetch(client, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("status = %+v, want %+v", got, tt.want)
			}
			wantCalls := 1
			if tt.indexed {
				wantCalls = 2
			}
			if calls != wantCalls {
				t.Errorf("requests = %d, want %d", calls, wantCalls)
			}
		})
	}
}

func TestAggregateFixtures(t *testing.T) {
	tests := []struct {
		file       string
		components map[string]any
	}{
		{"Shelly.GetConfig.json", map[string]any{
			"cloud":    CloudGetConfigResponse{Enable: true, Server: "shelly-0-eu.shelly.cloud:6022/jrpc"},
			"switch:0": SwitchGetConfigResponse{AutoOffDelay: 60, AutoOnDelay: 60, CurrentLimit: 10, InMode: "detached", InitialState: "on", Name: "switch1", PowerLimit: 700, VoltageLimit: 280},
			"switch:1": SwitchGetConfigResponse{Id: 1, AutoOffDelay: 60, AutoOnDelay: 60, CurrentLimit: 10, InMode: "detached", InitialState: "on", Name: "switch2", PowerLimit: 2800, VoltageLimit: 280},
		}},
		{"Shelly.GetStatus.json", map[string]any{
			"cloud":   CloudGetStatusResponse{Connected: true},
			"input:0": InputGetStatusResponse{},
			"input:1": InputGetStatusResponse{Id: 1},
			"wifi":    WifiGetStatusResponse{RSSI: -56, SSID: "Wifi 1", IP: "192.168.1.2", Status: "got ip"},
			"switch:0": SwitchGetStatusResponse{
				Source: "init", Output: true, Current: 0.162, Voltage: 229.5,
				ActiveEnergy: EnergyCounter{Total: 204.452, ByMinute: []float64{0, 0, 0}, MinuteTs: 1672814679},
				Temperature:  &Temperature{Celsius: 43.3, Fahrenheit: 109.9},
			},
			"switch:1": SwitchGetStatusResponse{
				Id: 1, Source: "init", Output: true, Current: 0.353, Voltage: 229.5, Power: 6.5, PowerFactor: -0.08,
				ActiveEnergy: EnergyCounter{Total: 277.009, ByMinute: []float64{76.27, 114.557, 114.602}, MinuteTs: 1672814679},
				Temperature:  &Temperature{Celsius: 43.3, Fahrenheit: 109.9},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			var components map[string]json.RawMessage
			if err := json.Unmarshal(body, &components); err != nil {
				t.Fatal(err)
			}
			for name, want := range tt.components {
				t.Run(name, func(t *testing.T) {
					got := reflect.New(reflect.TypeOf(want))
					if err := json.Unmarshal(components[name], got.Interface()); err != nil {
						t.Fatal(err)
					}
					if !reflect.DeepEqual(got.Elem().Interface(), want) {
						t.Errorf("component = %+v, want %+v", got.Elem().Interface(), want)
					}
				})
			}
		})
	}
}
