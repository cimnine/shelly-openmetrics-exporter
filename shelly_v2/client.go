package shelly_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dac "github.com/ndecker/go-http-digest-auth-client"
)

func (s *ShellyV2) FetchStatus() error {
	status := &Status{}

	err := s.getSwitchStatus(status)
	if err != nil {
		return err
	}
	err = s.getSwitchConfig(status)
	if err != nil {
		return err
	}
	err = s.getInputStatus(status)
	if err != nil {
		return err
	}
	err = s.getWifiStatus(status)
	if err != nil {
		return err
	}
	err = s.getCloudStatus(status)
	if err != nil {
		return err
	}
	err = s.getCloudConfig(status)
	if err != nil {
		return err
	}
	err = s.getVoltmeterStatus(status)
	if err != nil {
		return err
	}
	err = s.getTemperatureStatus(status)
	if err != nil {
		return err
	}
	err = s.getHumidityStatus(status)
	if err != nil {
		return err
	}
	err = s.getDevicePowerStatus(status)
	if err != nil {
		return err
	}
	err = s.getCoverStatus(status)
	if err != nil {
		return err
	}
	err = s.getCoverConfig(status)
	if err != nil {
		return err
	}
	err = s.getPM1Status(status)
	if err != nil {
		return err
	}

	s.status = status

	return nil
}

func (s *ShellyV2) do(request JsonRpc2Request, statusRes any) (end bool, err error) {
	res, err := s.sendRequest(request, err)
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
	code := resPayload.Error.Code
	if endLoop(code) {
		return true, nil
	}
	if resPayload.Error.Code != 0 {
		return true, fmt.Errorf("shelly returned error with code %d: %s", resPayload.Error.Code, resPayload.Error.Message)
	}

	return false, json.Unmarshal(resPayload.Result, statusRes)
}

func (s *ShellyV2) sendRequest(request JsonRpc2Request, err error) (*http.Response, error) {
	req, err := s.prepareHttpRequest(request)
	if err != nil {
		return nil, err
	}

	var res *http.Response
	if s.Username != "" && s.Password != "" {
		t := dac.NewTransport(s.Username, s.Password)
		t.Client = s.client
		res, err = t.RoundTrip(req)
	} else {
		res, err = s.client.Do(req)
	}

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"the shelly '%s' returned with an http status of %d, whereas only 200 was expected",
			s.TargetHost,
			res.StatusCode,
		)
	}

	return res, nil
}

func (s *ShellyV2) prepareHttpRequest(request JsonRpc2Request) (*http.Request, error) {
	request.MessageID = *s.nextMessageID

	reqPayload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	*s.nextMessageID = *s.nextMessageID + 1

	//goland:noinspection HttpUrlsUsage
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%s/rpc", s.TargetHost), bytes.NewBuffer(reqPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", s.UserAgent)
	return req, nil
}

func endLoop(code int) bool {
	const CodeIdNotFound = -105
	const CodeHandlerNotFound = 404

	return code == CodeIdNotFound ||
		code == CodeHandlerNotFound
}
