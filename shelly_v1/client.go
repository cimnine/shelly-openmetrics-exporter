package shelly_v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s *ShellyV1) FetchStatus() error {
	client := &http.Client{}

	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/status", s.targetHost), nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "shelly-prometheus-exporter")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.Body == nil {
		return fmt.Errorf("the response body was nil")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	s.status = &Status{}
	err = json.Unmarshal(body, s.status)

	return nil
}
