package core

import (
	"time"
)

type TimeRange struct {
	StartSec, EndSec int
}

func (r TimeRange) Start() string {
	return r.IntToString(r.StartSec)
}

func (r TimeRange) End() string {
	return r.IntToString(r.EndSec)
}

func (r TimeRange) IntToString(s int) string {
	t := time.Unix(int64(s), 0).UTC()
	return t.Format("15:04")
}
