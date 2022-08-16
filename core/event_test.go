package core

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEvent_Test(t *testing.T) {
	type fields struct {
		ID       uuid.UUID
		Name     string
		Schedule MapArrTimeRange
	}
	type args struct {
		params SlotParams
	}

	tests := []struct {
		name string
		fields
		args
		want []AvailableDay
	}{
		// {
		// 	name: "set 1 time range in September monday, should get all slot on September",
		// 	fields: fields{
		// 		Schedule: MapArrTimeRange{
		// 			time.Monday: []TimeRange{
		// 				{
		// 					StartSec: 0,
		// 					EndSec:   3600,
		// 				},
		// 			},
		// 		},
		// 	},
		// 	args: args{
		// 		params: SlotParams{
		// 			M: time.September,
		// 			Y: 2022,
		// 		},
		// 	},
		// 	want: []AvailableDay{
		// 		{
		// 			Date: time.Date(2022, time.September, 5, 0, 0, 0, 0, time.UTC),
		// 			Slot: []TimeRange{
		// 				{
		// 					StartSec: 0,
		// 					EndSec:   3600,
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Date: time.Date(2022, time.September, 12, 0, 0, 0, 0, time.UTC),
		// 			Slot: []TimeRange{
		// 				{
		// 					StartSec: 0,
		// 					EndSec:   3600,
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Date: time.Date(2022, time.September, 19, 0, 0, 0, 0, time.UTC),
		// 			Slot: []TimeRange{
		// 				{
		// 					StartSec: 0,
		// 					EndSec:   3600,
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Date: time.Date(2022, time.September, 26, 0, 0, 0, 0, time.UTC),
		// 			Slot: []TimeRange{
		// 				{
		// 					StartSec: 0,
		// 					EndSec:   3600,
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		{
			name: "set 2 time range in August sunday, should get all slot on August",
			fields: fields{
				Schedule: MapArrTimeRange{
					time.Sunday: []TimeRange{
						{
							StartSec: 0,
							EndSec:   1800,
						},
						{
							StartSec: 3600,
							EndSec:   7200,
						},
					},
				},
			},
			args: args{
				params: SlotParams{
					D: 7,
					M: time.August,
					Y: 2022,
				},
			},
			want: []AvailableDay{
				{
					Date: time.Date(2022, time.August, 7, 0, 0, 0, 0, time.UTC),
					Slot: []TimeRange{
						{
							StartSec: 0,
							EndSec:   1800,
						},
						{
							StartSec: 3600,
							EndSec:   7200,
						},
					},
				}, {
					Date: time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				}, {
					Date: time.Date(2022, time.August, 21, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				}, {
					Date: time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				},
			},
		},
		{
			name: "set 1 time range in September on monday without spesified day, should get all available day on September with nil slot",
			fields: fields{
				Schedule: MapArrTimeRange{
					time.Monday: []TimeRange{
						{
							StartSec: 0,
							EndSec:   3600,
						},
					},
				},
			},
			args: args{
				params: SlotParams{
					M: time.September,
					Y: 2022,
				},
			},
			want: []AvailableDay{
				{
					Date: time.Date(2022, time.September, 5, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				},
				{
					Date: time.Date(2022, time.September, 12, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				},
				{
					Date: time.Date(2022, time.September, 19, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				},
				{
					Date: time.Date(2022, time.September, 26, 0, 0, 0, 0, time.UTC),
					Slot: nil,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Event{
				Schedules: tt.fields.Schedule,
			}
			res, _ := r.GetAvailableSlots(tt.args.params)

			assert.Equal(t, tt.want, res)
		})
	}
}
