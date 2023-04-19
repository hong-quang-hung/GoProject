package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func Make[T any](size int, init T) []T {
	slice := make([]T, size)
	for i := range slice {
		slice[i] = init
	}
	return slice
}

func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Slice[T any](a ...T) []T {
	return a
}

func Insert[T any](a []T, index int, value T) []T {
	if index >= len(a) {
		a = append(a, value)
	} else {
		a = append(a[:index+1], a[index:]...)
		a[index] = value
	}
	return a
}

func InsertInt(a []int, index int, value int) []int {
	if index >= len(a) {
		a = append(a, value)
	} else {
		a = append(a[:index+1], a[index:]...)
		a[index] = value
	}
	return a
}

func S2SliceString(s string) []string {
	re := regexp.MustCompile(`^\[(.+)\]$`)
	matched := re.FindAllStringSubmatch(s, -1)
	if len(matched) == 0 || len(strings.TrimSpace(matched[0][1])) == 0 {
		return []string{}
	}

	arr := strings.Split(matched[0][1], ",")
	res := make([]string, len(arr))
	for i, a := range arr {
		a = strings.Trim(a, " ")
		res[i] = a
	}
	return res
}

func S2SliceInt(s string) []int {
	re := regexp.MustCompile(`^\[(.+)\]$`)
	matched := re.FindAllStringSubmatch(s, -1)
	if len(matched) == 0 || len(strings.TrimSpace(matched[0][1])) == 0 {
		return []int{}
	}

	arr := strings.Split(matched[0][1], ",")
	res := make([]int, len(arr))
	for i, a := range arr {
		a = strings.Trim(a, " ")
		val, _ := strconv.ParseInt(a, 10, 0)
		res[i] = int(val)
	}
	return res
}

func S2SoSliceInt(s string) [][]int {
	s = strings.Trim(s, " ")
	s = s[1 : len(s)-1]
	re := regexp.MustCompile(`(\[[^\[]*\])`)
	arr := re.FindAllStringSubmatch(s, -1)
	res := make([][]int, len(arr))
	for i, a := range arr {
		res[i] = S2SliceInt(a[1])
	}
	return res
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
