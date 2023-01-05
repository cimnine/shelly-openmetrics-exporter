package shelly

type Shelly struct {
	TargetHost string
	UserAgent  string
	Username   string
	Password   string
}

type Device interface {
	FetchStatus() error
	FillMetrics(m *Metrics)
}
