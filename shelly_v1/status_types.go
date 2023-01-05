package shelly_v1

type Emeter struct {
	Current       float64 `json:"current"`
	IsValid       bool    `json:"is_valid"`
	PowerFactor   float64 `json:"pf"`
	Power         float64 `json:"power"`
	Total         float64 `json:"total"`
	TotalReturned float64 `json:"total_returned"`
	Voltage       float64 `json:"voltage"`
}

type Relay struct {
	HasTimer       bool   `json:"has_timer"`
	IsValid        bool   `json:"is_valid"`
	IsOn           bool   `json:"ison"`
	Overpower      bool   `json:"overpower"`
	Source         string `json:"source"`
	TimerDuration  int    `json:"timer_duration"`
	TimerRemaining int    `json:"timer_remaining"`
	TimerStarted   int    `json:"timer_started"`
}

type UpdateStatus struct {
	HasUpdate  bool   `json:"has_update"`
	NewVersion string `json:"new_version"`
	OldVersion string `json:"old_version"`
	Status     string `json:"status"`
}

type WifiStatus struct {
	Connected bool   `json:"connected"`
	Ip        string `json:"ip"`
	Rssi      int    `json:"rssi"`
	Ssid      string `json:"ssid"`
}

type MQTTStatus struct {
	Connected bool `json:"connected"`
}

type EmeterN struct {
	Current  float64 `json:"current"`
	IsValid  bool    `json:"is_valid"`
	Ixsum    float64 `json:"ixsum"`
	Mismatch bool    `json:"mismatch"`
}

type CloudStatus struct {
	Connected bool `json:"connected"`
	Enabled   bool `json:"enabled"`
}

type ActionStats struct {
	Skipped int `json:"skipped"`
}

type Temperature struct {
	Celsius    float64 `json:"tC"`
	Fahrenheit float64 `json:"tF"`
	IsValid    bool    `json:"is_valid"`
}

type ADC struct {
	Voltage float64 `json:"voltage"`
}

type Input struct {
	Input      int    `json:"input"`
	Event      string `json:"event"`
	EventCount int    `json:"event_cnt"`
}

type Status struct {
	ActionsStats      *ActionStats  `json:"actions_stats"`
	ADCs              []ADC         `json:"adcs"`
	CfgChangedCnt     int           `json:"cfg_changed_cnt"`
	Cloud             *CloudStatus  `json:"cloud"`
	CtCalst           int           `json:"ct_calst"`
	EmeterN           *EmeterN      `json:"emeter_n"`
	Emeters           []Emeter      `json:"emeters"`
	FsFree            int           `json:"fs_free"`
	FsMounted         bool          `json:"fs_mounted"`
	FsSize            int           `json:"fs_size"`
	HasUpdate         bool          `json:"has_update"`
	Inputs            []Input       `json:"inputs"`
	Mac               string        `json:"mac"`
	Mqtt              *MQTTStatus   `json:"mqtt"`
	RamFree           int           `json:"ram_free"`
	RamTotal          int           `json:"ram_total"`
	Relays            []Relay       `json:"relays"`
	Serial            int           `json:"serial"`
	Time              string        `json:"time"`
	TotalPower        float64       `json:"total_power"`
	Temperature       float64       `json:"temperature"`
	Overtemperature   bool          `json:"overtemperature"`
	Temperatures      *Temperature  `json:"tmp"`
	TemperatureStatus string        `json:"temperature_status"`
	Unixtime          int           `json:"unixtime"`
	Update            *UpdateStatus `json:"update"`
	Uptime            int           `json:"uptime"`
	VData             int           `json:"v_data"`
	WifiSta           *WifiStatus   `json:"wifi_sta"`
}
