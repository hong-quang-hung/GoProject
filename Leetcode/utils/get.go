package utils

func Iff[T any](c bool, a T, b T) T {
	if c {
		return a
	}
	return b
}

func IffNil[T any](c any, a T, b T) T {
	if c == nil {
		return a
	}
	return b
}
