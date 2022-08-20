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
		want []time.Time
	}{
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
					star: time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
					end:  time.Date(2022, time.September, 0, 0, 0, 0, 0, time.UTC),
				},
			},
			want: []time.Time{
				time.Date(2022, time.August, 7, 0, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 21, 0, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 28, 0, 0, 0, 0, time.UTC),
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

func TestGetAvailableTimeStart_Test(t *testing.T) {
	type fields struct {
		ID       uuid.UUID
		Name     string
		Schedule MapArrTimeRange
		Duration int
	}
	type args struct {
		day time.Time
	}

	tests := []struct {
		name string
		fields
		args
		want []time.Time
	}{
		{
			name: "set 2 time range in 14 August sunday, should get all available time start slot on that day",
			fields: fields{
				Schedule: MapArrTimeRange{
					time.Sunday: []TimeRange{
						{
							StartSec: 0,
							EndSec:   3600,
						},
						{
							StartSec: 7200,
							EndSec:   10800,
						},
					},
				},
				Duration: 30,
			},
			args: args{
				day: time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
			},
			want: []time.Time{
				time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 0, 30, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 2, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 2, 30, 0, 0, time.UTC),
			},
		},
		{
			name: "set 2 time range in 14 August sunday, should get all available time start slot on that day",
			fields: fields{
				Schedule: MapArrTimeRange{
					time.Sunday: []TimeRange{
						{
							StartSec: 0,
							EndSec:   3600,
						},
						{
							StartSec: 3600,
							EndSec:   7200,
						},
					},
				},
				Duration: 30,
			},
			args: args{
				day: time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
			},
			want: []time.Time{
				time.Date(2022, time.August, 14, 0, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 0, 30, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 1, 0, 0, 0, time.UTC),
				time.Date(2022, time.August, 14, 1, 30, 0, 0, time.UTC),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Event{
				Schedules: tt.fields.Schedule,
				Duration:  tt.Duration,
			}

			res, _ := r.GetAllSlots(tt.args.day)

			assert.Equal(t, tt.want, res)
		})
	}
}
