package main

import (
	"testing"
	"time"

	"github.com/GnvSaikiran/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 4, 3, 10, 30, 0, 0, time.UTC),
			want: "03 Apr 2024 at 10:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 4, 3, 10, 30, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "03 Apr 2024 at 09:30",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := humanDate(test.tm)

			assert.Equal(t, test.want, got)
		})
	}
}
