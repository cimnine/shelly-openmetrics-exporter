package shelly_detect

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestDetectionFixtures(t *testing.T) {
	tests := []struct {
		file       string
		generation ShellyGeneration
		want       Shelly
	}{
		{"shelly_v1.json", ShellyGeneration1, Shelly{
			Type: "SHEM-3", Mac: "111CAB11A11E", Fw: "20221027-110030/v1.12.1-ga9117d3",
			LongID: 1, NumEmeters: 3, NumOutputs: 1, ReportPeriod: 1,
		}},
		{"shelly_v2.json", ShellyGeneration2, Shelly{
			Id: "shellyplus2pm-11111CED111", Mac: "11111CED111", Model: "SNSW-102P16EU", Gen: 2,
			FwId: "20221206-141957/0.12.0-gafc2404", Ver: "0.12.0", App: "Plus2PM",
			Name: "shellyplus2pm-11111CED111", Profile: "switch",
		}},
		{"../shelly_v2/Shelly.GetDeviceInfo.json", ShellyGeneration2, Shelly{
			Id: "shellyplus2pm-11111CED111", Mac: "11111CED111", Model: "SNSW-102P16EU", Gen: 2,
			FwId: "20221206-141957/0.12.0-gafc2404", Ver: "0.12.0", App: "Plus2PM",
			Name: "shellyplus2pm-11111CED111", Profile: "switch",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			body, err := os.ReadFile(tt.file)
			if err != nil {
				t.Fatal(err)
			}
			var got Shelly
			if err := json.Unmarshal(body, &got); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decoded device = %+v, want %+v", got, tt.want)
			}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/shelly" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write(body)
			}))
			defer server.Close()
			generation, err := DetectVersion(strings.TrimPrefix(server.URL, "http://"))
			if err != nil {
				t.Fatal(err)
			}
			if generation != tt.generation {
				t.Errorf("generation = %d, want %d", generation, tt.generation)
			}
		})
	}
}
