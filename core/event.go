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
	M time.Month
	Y int
}

func (e Event) GetAvailableSlots(params SlotParams) ([]AvailableDay, error) {
	star := time.Date(params.Y, params.M, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(params.Y, params.M+1, 0, 0, 0, 0, 0, time.UTC)

	var res []AvailableDay

	starDay := star.Day()
	endDay := end.Day()

	for i := starDay; i <= endDay; i++ {
		cur := time.Date(params.Y, params.M, i, 0, 0, 0, 0, time.UTC)
		if _, ok := e.Schedules[cur.Weekday()]; ok {
			res = append(res, AvailableDay{
				Date: cur,
				Slot: nil,
			})
		}
	}

	return res, nil
}

type Slot struct {
	Star int
}

type AvailableDay struct {
	Date time.Time
	Slot []Slot
}
