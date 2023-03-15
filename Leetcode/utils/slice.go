package utils

import (
	"leetcode.com/Leetcode/types"
)

func Remove[T types.Obj](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
