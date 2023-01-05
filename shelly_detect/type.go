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
	LongID       int    `json:"longid"`
	Discoverable bool   `json:"discoverable"`
	Name         string `json:"name"`
	Profile      string `json:"profile"`
	NumEmeters   int    `json:"num_emeters"`
	NumMeters    int    `json:"num_meters"`
	NumOutputs   int    `json:"num_outputs"`
	ReportPeriod int    `json:"report_period"`
}
