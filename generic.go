package helpers

func Ref[T any](in T) *T {
	return &in
}

func Deref[T comparable](in *T) T {
	if in == nil {
		var empty T
		return empty
	}
	return *in
}

func Overlaps[T comparable](a, b []T) bool {
	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				return true
			}
		}
	}
	return false
}
