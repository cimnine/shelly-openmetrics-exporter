package shelly_v2

import (
	"encoding/json"
)

type Auth struct {
	Realm     string `json:"realm"`
	Username  string `json:"username"`
	Nonce     int64  `json:"nonce"`
	Cnonce    int    `json:"cnonce"`
	Response  string `json:"response"`
	Algorithm string `json:"algorithm"`
}

type JsonRpc2Request struct {
	JsonRpcVersion string      `json:"jsonrpc,omitempty"`
	MessageID      int         `json:"id"`
	Src            string      `json:"src"`
	Method         string      `json:"method"`
	Params         interface{} `json:"params,omitempty"`
	Auth           *Auth       `json:"auth,omitempty"`
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
