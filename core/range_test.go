package core_test

import (
	"testing"

	"github.com/rudikurniawan99/calendly-go.git/core"
	"github.com/stretchr/testify/assert"
)

func TestTimeRange_Test(t *testing.T) {
	type fields struct {
		StartSec, EndSec int
	}

	test := []struct {
		name      string
		fields    fields
		wantStart string
		wantEnd   string
	}{
		{
			name: "1 oclock",
			fields: fields{
				StartSec: 3600,
				EndSec:   7200,
			},
			wantStart: "01:00",
			wantEnd:   "02:00",
		}, {
			name: "2 oclock",
			fields: fields{
				StartSec: 3600,
				EndSec:   14400,
			},
			wantStart: "01:00",
			wantEnd:   "04:00",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			r := core.TimeRange{
				StartSec: tt.fields.StartSec,
				EndSec:   tt.fields.EndSec,
			}

			assert.Equal(t, tt.wantStart, r.Start())
			assert.Equal(t, tt.wantEnd, r.End())
		})
	}
}
