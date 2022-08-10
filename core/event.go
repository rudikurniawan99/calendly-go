package core

import (
	"time"

	"github.com/google/uuid"
)

//schedule store information about availability for each day
type Event struct {
	ID        uuid.UUID
	Name      string
	Schedules map[time.Weekday][]TimeRange
}

type SlotParams struct {
	Month time.Month
	Day   *time.Time
}

func (e Event) GetAvailableSlots(params SlotParams) ([]AvailableDay, error) {
	return nil, nil
}

type Slot struct {
	Star int
}

type AvailableDay struct {
	Slot []Slot
}
