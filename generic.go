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
