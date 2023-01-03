package shelly_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchStatus(targetHost string) (status Status, err error) {
	for i := 0; true; i++ {
		statusRes := SwitchGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			MessageId:      i,
			Src:            "shelly-prometheus-exporter",
			Method:         "Switch.GetStatus",
			Params:         SwitchGetStatusRequest{Id: i},
		}

		end, err2 := do(targetHost, request, &statusRes)
		err = err2
		if end {
			break
		}
		if err != nil {
			return
		}

		status.Switches = append(status.Switches, statusRes)
	}

	return
}

func do(targetHost string, request JsonRpc2Request, statusRes *SwitchGetStatusResponse) (end bool, err error) {
	client := &http.Client{}

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return
	}

	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/rpc", targetHost), bytes.NewBuffer(reqPayload))
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "shelly-prometheus-exporter")

	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.Body == nil {
		return false, fmt.Errorf("the response body was nil")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	resPayload := JsonRpc2Response{}
	err = json.Unmarshal(body, &resPayload)
	if err != nil {
		return
	}
	if resPayload.Error.Code == -105 { // Switch ID not found
		return true, nil
	}
	if resPayload.Error.Code != 0 {
		return false, fmt.Errorf("shelly returned error with code %d: %s", resPayload.Error.Code, resPayload.Error.Message)
	}

	return false, json.Unmarshal(resPayload.Result, statusRes)
}
