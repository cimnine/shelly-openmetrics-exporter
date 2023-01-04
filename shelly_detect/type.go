package shelly_detect

type Shelly struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Mac          string `json:"mac"`
	Model        string `json:"model"`
	Gen          int    `json:"gen"`
	Fw           string `json:"fw"`
	FwId         string `json:"fw_id"`
	Ver          string `json:"ver"`
	App          string `json:"app"`
	Auth         bool   `json:"auth"`
	AuthEn       bool   `json:"auth_en"`
	AuthDomain   string `json:"auth_domain"`
	Longid       int    `json:"longid"`
	Discoverable bool   `json:"discoverable"`
}
