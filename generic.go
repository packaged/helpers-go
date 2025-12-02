package helpers

import "reflect"

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

// If A generic ternary function which returns vtrue if cond is true, otherwise it returns vfalse.
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func Coalesce[T any](input ...T) T {
	var check T
	for _, check = range input {
		if !IsZero(check) {
			return check
		}
	}
	return check
}

func IsZero(input any) bool {
	if input == nil {
		return true
	}
	r := reflect.ValueOf(input)
	return !r.IsValid() || (r.Kind() == reflect.Pointer && r.IsNil()) || r.IsZero()
}
