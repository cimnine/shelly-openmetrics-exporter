package shelly_detect

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ShellyGeneration int

const ShellyGeneration1 ShellyGeneration = 1
const ShellyGeneration2 ShellyGeneration = 2

func DetectVersion(targetHost string) (ShellyGeneration, error) {
	client := &http.Client{}

	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/shelly", targetHost), nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "shelly-prometheus-exporter")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	if res.Body == nil {
		err = fmt.Errorf("the response body was nil")
		return 0, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	shelly := Shelly{}
	err = json.Unmarshal(body, &shelly)

	var gen ShellyGeneration
	if shelly.Gen == 0 {
		gen = ShellyGeneration1
	} else {
		gen = ShellyGeneration(shelly.Gen)
	}

	return gen, nil
}
