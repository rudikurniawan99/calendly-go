package core

import (
	"time"

	"github.com/google/uuid"
)

type MapArrTimeRange map[time.Weekday][]TimeRange

//schedule store information about availability for each day
type Event struct {
	ID        uuid.UUID
	Name      string
	Schedules MapArrTimeRange
}

type SlotParams struct {
	M    time.Month
	D, Y int
}

func (e Event) GetAvailableSlots(params SlotParams) ([]AvailableDay, error) {
	star := time.Date(params.Y, params.M, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(params.Y, params.M+1, 0, 0, 0, 0, 0, time.UTC)

	var res []AvailableDay

	starDay := star.Day()
	endDay := end.Day()

	for i := starDay; i <= endDay; i++ {
		cur := time.Date(params.Y, params.M, i, 0, 0, 0, 0, time.UTC)
		if rs, ok := e.Schedules[cur.Weekday()]; ok {

			if params.D != cur.Day() {
				rs = nil
			}

			res = append(res, AvailableDay{
				Date: cur,
				Slot: rs,
			})
		}
	}

	return res, nil
}

func (e Event) GetTimeStart() string {
	// e.Schedules[]
	return ""
}

type AvailableDay struct {
	Date time.Time
	Slot []TimeRange
}
