package core

import "strconv"

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
	var strHour, strMinute string
	hour := s / 3600
	minutes := (s % 3600) / 60

	if hour == 24 {
		strHour = "00"
	} else {
		strHour = convert(hour)
	}
	strMinute = convert(minutes)

	return strHour + ":" + strMinute
}

func convert(t int) string {
	if t < 10 {
		return "0" + strconv.Itoa(t)
	}
	return strconv.Itoa(t)
}
