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
	Duration  int
	Location  time.Location
}

type SlotParams struct {
	star, end time.Time
}

func (e Event) GetAvailableSlots(params SlotParams) ([]time.Time, error) {
	var res []time.Time

	starDay := params.star
	endDay := params.end

	cur := starDay
	for i := starDay.Day(); i <= endDay.Day(); i++ {
		if _, ok := e.Schedules[cur.Weekday()]; ok {
			res = append(res, cur)
		}
		cur = cur.AddDate(0, 0, 1)
	}

	return res, nil
}

func (e Event) GetAllSlots(day time.Time) ([]time.Time, error) {
	var res []time.Time
	timeRange := e.Schedules[day.Weekday()]
	duration := time.Duration(e.Duration) * time.Minute

	for _, t := range timeRange {
		cur := day.Add(time.Duration(t.StartSec) * time.Second)
		end := day.Add(time.Duration(t.EndSec) * time.Second)

		for cur.Before(end) {
			endTime := cur.Add(duration)

			if endTime.Before(end) || endTime == end {
				res = append(res, cur)
			}
			cur = endTime
		}
	}
	return res, nil
}

// func (e Event) GetAllSlots(starTime, endTime time.Time) ([]time.Time, error) {
// 	var availableSlots []time.Time
// 	duration := time.Duration(e.Duration) * time.Minute

// 	cur := starTime

// 	for {
// 		if cur == endTime || cur.After(endTime) {
// 			break
// 		}

// 		end := cur.Add(duration)
// 		if end == endTime || end.Before(endTime) {
// 			availableSlots = append(availableSlots, cur)
// 		}
// 		cur = end
// 	}

// 	return availableSlots, nil
// }
