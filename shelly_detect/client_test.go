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
			Type: "SHEM-3", Mac: "020000000001", Fw: "20230913-114244/v1.14.0-gcb84623",
			LongID: 1, NumEmeters: 3, NumOutputs: 1, ReportPeriod: 1,
		}},
		{"shelly_v2.json", ShellyGeneration2, Shelly{
			Id: "shellyplus2pm-020000000003", Mac: "020000000003", Model: "SNSW-102P16EU", Gen: 2,
			FwId: "20260311-095847/1.7.5-g9979d16", Ver: "1.7.5", App: "Plus2PM",
			Name: "Switch device", Profile: "switch",
		}},
		{"../shelly_v2/Shelly.GetDeviceInfo.json", ShellyGeneration2, Shelly{
			Id: "shellyplus2pm-020000000003", Mac: "020000000003", Model: "SNSW-102P16EU", Gen: 2,
			FwId: "20260311-095847/1.7.5-g9979d16", Ver: "1.7.5", App: "Plus2PM",
			Name: "Switch device", Profile: "switch",
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
