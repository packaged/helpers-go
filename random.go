package helpers

import "math/rand/v2"

func RandomItem[T any](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}

	index := rand.IntN(len(values))
	return values[index]
}
