package utils

import (
	"leetcode.com/Leetcode/types"
)

func Max[T types.Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T types.Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
