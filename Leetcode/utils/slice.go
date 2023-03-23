package utils

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Slice[T any](a ...T) []T {
	return a
}

func SliceInt(s string) []int {
	return nil
}
