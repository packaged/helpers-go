package helpers

import (
	"math"
	"time"
)

func IntToUint32(i int) uint32 {
	if i < 0 {
		return 0
	}
	if i > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(i)
}

func IntToInt32(i int) int32 {
	return Int64ToInt32(int64(i))
}

func MonthToInt32(i time.Month) int32 {
	if i < time.January {
		i = time.January
	}
	if i > time.December {
		i = time.December
	}
	return IntToInt32(int(i))
}

func Int64ToInt32(i int64) int32 {
	if i < math.MinInt32 {
		return math.MinInt32
	}
	if i > math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(i)
}
