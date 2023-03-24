package utils

import "strings"

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Slice[T any](a ...T) []T {
	return a
}

func SliceInt(s string) []int {
	return nil
}

func Sslice[T Number](slice []T) string {
	sep := ","
	s, _ := joinSlice(slice, sep)
	return s
}

func SsoSlice[T Number](soSlice [][]T) string {
	var lenght int = len(soSlice)
	switch lenght {
	case 0:
		return ""
	case 1:
		return Sslice(soSlice[0])
	}

	sep := ","
	str := make([]string, lenght)
	n := len(sep) * (lenght - 1)
	for i := 0; i < lenght; i++ {
		s, l := joinSlice(soSlice[i], sep)
		str[i] = s
		n += l + 2
	}

	var b strings.Builder
	b.Grow(n + 2)
	b.WriteString("[")
	b.WriteString(str[0])
	for _, s := range str[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	b.WriteString("]")
	return b.String()
}

func joinSlice[T Number](slice []T, sep string) (string, int) {
	s, l := join(slice, sep)
	var b strings.Builder
	b.Grow(l + 2)
	b.WriteString("[")
	b.WriteString(s)
	b.WriteString("]")
	return b.String(), b.Len()
}

func join[T Number](slice []T, sep string) (string, int) {
	switch len(slice) {
	case 0:
		return "", 0
	case 1:
		return ToString(slice[0]), Len(slice[0])
	}
	n := len(sep) * (len(slice) - 1)
	for i := 0; i < len(slice); i++ {
		n += Len(slice[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(ToString(slice[0]))
	for _, s := range slice[1:] {
		b.WriteString(sep)
		b.WriteString(ToString(s))
	}
	return b.String(), b.Len()
}
