package shelly_v2

import (
	"encoding/json"
)

type JsonRpc2Request struct {
	JsonRpcVersion string      `json:"jsonrpc"`
	MessageID      int         `json:"id"`
	Src            string      `json:"src"`
	Method         string      `json:"method"`
	Params         interface{} `json:"params"`
}

type JsonRpc2Response struct {
	MessageId int             `json:"id"`
	Src       string          `json:"src"`
	Dst       string          `json:"dst"`
	Result    json.RawMessage `json:"result"`
	Error     struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
}
