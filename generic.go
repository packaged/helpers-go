package helpers

// Ref is deprecated.
//
//go:fix inline
func Ref[T any](in T) *T {
	return new(in)
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

// If A generic ternary function which returns vtrue if cond is true, otherwise it returns vfalse.
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}
