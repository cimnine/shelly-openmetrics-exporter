package shelly

type Shelly struct {
	TargetHost string
}

type Actor interface {
	FetchStatus() error
	FillMetrics(m *Metrics)
}

const CelsiusInKelvin = 273.15
