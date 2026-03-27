package helpers

import "time"

// IsZeroTime returns true if the given time is either the zero value or Unix epoch (1970-01-01T00:00:00Z).
func IsZeroTime(t time.Time) bool {
	return t.IsZero() || t.Unix() == 0
}