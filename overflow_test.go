package helpers

import (
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntToInt32(t *testing.T) {
	tests := []struct {
		in           int
		expected     int32
		expectedFail int
	}{
		{0, 0, 0},
		{math.MinInt32, math.MinInt32, math.MinInt32},
		{math.MinInt32 - 1, math.MinInt32, math.MaxInt32}, // overflow becomes max
		{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		{math.MaxInt32 + 1, math.MaxInt32, math.MinInt32}, // overflow becomes min
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			assert.EqualValues(t, test.expected, IntToInt32(test.in))
			assert.EqualValues(t, test.expectedFail, int32(test.in)) // #nosec G115 -- intentional overflow to test safe conversion
		})
	}
}

func TestIntToUint32(t *testing.T) {
	tests := []struct {
		in           int
		expected     uint32
		expectedFail int
	}{
		{0, 0, 0},
		{math.MinInt, 0, 0},
		{math.MinInt32 - 1, 0, math.MaxInt32},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		{math.MaxInt32 + 1, math.MaxInt32 + 1, math.MaxInt32 + 1},
		{math.MaxUint32 + 1, math.MaxUint32, 0}, // overflow becomes zero
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.in), func(t *testing.T) {
			assert.EqualValues(t, test.expected, IntToUint32(test.in))
			assert.EqualValues(t, test.expectedFail, uint32(test.in)) // #nosec G115 -- intentional overflow to test safe conversion
		})
	}
}

func TestMonthToInt32(t *testing.T) {
	tests := []struct {
		in       time.Month
		expected int32
	}{
		{time.January - 1, 1},
		{time.January, 1},
		{time.February, 2},
		{time.March, 3},
		{time.April, 4},
		{time.May, 5},
		{time.June, 6},
		{time.July, 7},
		{time.August, 8},
		{time.September, 9},
		{time.October, 10},
		{time.November, 11},
		{time.December, 12},
		{time.December + 1, 12},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(int(test.in)), func(t *testing.T) {
			assert.EqualValues(t, test.expected, MonthToInt32(test.in))
		})
	}
}
