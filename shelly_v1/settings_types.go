package shelly_v1

import (
	"time"
)

type Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

type APRoaming struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
}

type BuildInfo struct {
	BuildId        string    `json:"build_id"`
	BuildTimestamp time.Time `json:"build_timestamp"`
	BuildVersion   string    `json:"build_version"`
}

type CloudSettings struct {
	Connected bool `json:"connected"`
	Enabled   bool `json:"enabled"`
}

type CoIOT struct {
	Enabled      bool   `json:"enabled"`
	Peer         string `json:"peer"`
	UpdatePeriod int    `json:"update_period"`
}

type Device struct {
	Hostname     string `json:"hostname"`
	Mac          string `json:"mac"`
	NumEmeters   int    `json:"num_emeters"`
	NumMeters    int    `json:"num_meters"`
	NumOutputs   int    `json:"num_outputs"`
	ReportPeriod int    `json:"report_period"`
	Type         string `json:"type"`
}

type EmeterNSettings struct {
	MismatchThreshold float64 `json:"mismatch_threshold"`
	RangeExtender     int     `json:"range_extender"`
}

type EmeterSettings struct {
	ApplianceType string `json:"appliance_type"`
	MaxPower      int    `json:"max_power"`
	Name          string `json:"name"`
	RangeExtender int    `json:"range_extender"`
}

type HardwareInfo struct {
	BatchId    int    `json:"batch_id"`
	HwRevision string `json:"hw_revision"`
}

type LoginSettings struct {
	Enabled     bool   `json:"enabled"`
	Unprotected bool   `json:"unprotected"`
	Username    string `json:"username"`
}

type MQTTSettings struct {
	CleanSession        bool    `json:"clean_session"`
	Enable              bool    `json:"enable"`
	Id                  string  `json:"id"`
	KeepAlive           int     `json:"keep_alive"`
	MaxQos              int     `json:"max_qos"`
	ReconnectTimeoutMax float64 `json:"reconnect_timeout_max"`
	ReconnectTimeoutMin float64 `json:"reconnect_timeout_min"`
	Retain              bool    `json:"retain"`
	Server              string  `json:"server"`
	UpdatePeriod        int     `json:"update_period"`
	User                string  `json:"user"`
}

type RelaySettings struct {
	AutoOff       float64       `json:"auto_off"`
	AutoOn        float64       `json:"auto_on"`
	DefaultState  string        `json:"default_state"`
	HasTimer      bool          `json:"has_timer"`
	Ison          bool          `json:"ison"`
	Name          interface{}   `json:"name"`
	Schedule      bool          `json:"schedule"`
	ScheduleRules []interface{} `json:"schedule_rules"`
}

type SNTPSettings struct {
	Enabled bool   `json:"enabled"`
	Server  string `json:"server"`
}

type WifiAPSettings struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
	Ssid    string `json:"ssid"`
}

type WifiSettings struct {
	Dns        interface{} `json:"dns"`
	Enabled    bool        `json:"enabled"`
	Gw         interface{} `json:"gw"`
	Ip         interface{} `json:"ip"`
	Ipv4Method string      `json:"ipv4_method"`
	Mask       interface{} `json:"mask"`
	Ssid       string      `json:"ssid"`
}

type Settings struct {
	Actions          Actions          `json:"actions"`
	AllowCrossOrigin bool             `json:"allow_cross_origin"`
	ApRoaming        APRoaming        `json:"ap_roaming"`
	BuildInfo        BuildInfo        `json:"build_info"`
	CfOutput         int              `json:"cf_output"`
	Cloud            CloudSettings    `json:"cloud"`
	Coiot            CoIOT            `json:"coiot"`
	DebugEnable      bool             `json:"debug_enable"`
	Device           Device           `json:"device"`
	Discoverable     bool             `json:"discoverable"`
	EcoModeEnabled   bool             `json:"eco_mode_enabled"`
	EmeterN          EmeterNSettings  `json:"emeter_n"`
	Emeters          []EmeterSettings `json:"emeters"`
	Firmware         string           `json:"fw"`
	Hwinfo           HardwareInfo     `json:"hwinfo"`
	Lat              float64          `json:"lat"`
	LedStatusDisable bool             `json:"led_status_disable"`
	Lng              float64          `json:"lng"`
	Login            LoginSettings    `json:"login"`
	Mqtt             MQTTSettings     `json:"mqtt"`
	Name             string           `json:"name"`
	PinCode          string           `json:"pin_code"`
	Relays           []RelaySettings  `json:"relays"`
	Sntp             SNTPSettings     `json:"sntp"`
	Time             string           `json:"time"`
	Timezone         string           `json:"timezone"`
	TzDst            bool             `json:"tz_dst"`
	TzDstAuto        bool             `json:"tz_dst_auto"`
	TzUtcOffset      int              `json:"tz_utc_offset"`
	Tzautodetect     bool             `json:"tzautodetect"`
	Unixtime         int              `json:"unixtime"`
	WifiAp           WifiAPSettings   `json:"wifi_ap"`
	WifiSta          WifiSettings     `json:"wifi_sta"`
	WifiSta1         WifiSettings     `json:"wifi_sta1"`
}
