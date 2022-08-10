package core

type TimeRange struct {
	StartSec, EndSec int
}

func (r TimeRange) Start() string {
	return r.Translate(r.StartSec)
}

func (r TimeRange) End() string {
	return r.Translate(r.EndSec)
}

// implementation later
func (r TimeRange) Translate(s int) string {
	return "02:00"
}
