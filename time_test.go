package helpers

import (
	"testing"
	"time"
)

func TestIsZeroTime(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want bool
	}{
		{"zero value", time.Time{}, true},
		{"unix epoch", time.Unix(0, 0), true},
		{"unix epoch UTC", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{"non-zero time", time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC), false},
		{"one second after epoch", time.Unix(1, 0), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZeroTime(tt.t); got != tt.want {
				t.Errorf("IsZeroTime() = %v, want %v", got, tt.want)
			}
		})
	}
}