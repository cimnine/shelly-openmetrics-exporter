package shelly_v2

type SwitchGetStatusRequest struct {
	Id int `json:"id"`
}

type EnergyCounter struct {
	Total    float64   `json:"total"`
	ByMinute []float64 `json:"by_minute"`
	MinuteTs int       `json:"minute_ts"`
}

type Temperature struct {
	Celsius    float64 `json:"tC"`
	Fahrenheit float64 `json:"tF"`
}

type SwitchGetStatusResponse struct {
	Id           int           `json:"id"`
	Source       string        `json:"source"`
	Output       bool          `json:"output"`
	Power        float64       `json:"apower"`
	PowerFactor  float64       `json:"pf"`
	Voltage      float64       `json:"voltage"`
	Current      float64       `json:"current"`
	ActiveEnergy EnergyCounter `json:"aenergy"`
	Temperature  *Temperature  `json:"temperature"`
}
