package shelly_v1

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchStatus(targetHost string) (status Status, err error) {
	client := &http.Client{}

	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/status", targetHost), nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "shelly-prometheus-exporter")

	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.Body == nil {
		err = fmt.Errorf("the response body was nil")
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &status)
	return
}
