package gg

// Map convert a slice to another by mapping each element
func Map[T any, U any](s []T, fn func(T) U) []U {
	var o []U
	if s == nil {
		return o
	}
	o = make([]U, len(s), len(s))
	for i := range s {
		o[i] = fn(s[i])
	}
	return o
}

// Set convert a slice of comparables into a Set-like map
func Set[T comparable](s []T) map[T]struct{} {
	o := make(map[T]struct{})
	for _, k := range s {
		o[k] = struct{}{}
	}
	return o
}

// Keys extract keys of map
func Keys[T comparable, U any](m map[T]U) []T {
	var o []T
	for k := range m {
		o = append(o, k)
	}
	return o
}

// Values extract values of map
func Values[T comparable, U any](m map[T]U) []U {
	var o []U
	for _, v := range m {
		o = append(o, v)
	}
	return o
}

// Filter returns a new slice with element filtered
func Filter[T any](s []T, fn func(T) bool) []T {
	var o []T
	for _, v := range s {
		if fn(v) {
			o = append(o, v)
		}
	}
	return o
}

// Repeat create a slice with a value repeated n times
func Repeat[T any](n int, v T) []T {
	out := make([]T, n, n)
	for i := range out {
		out[i] = v
	}
	return out
}
