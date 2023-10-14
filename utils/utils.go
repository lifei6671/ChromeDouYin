package utils

func First[T any](s []T) T {
	var zero T
	if len(s) == 0 {
		return zero
	}
	return s[0]
}
