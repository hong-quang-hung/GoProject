package utils

import (
	"strconv"

	"leetcode.com/Leetcode/types"
)

const (
	commaSpaceString = ","
	nilStr           = "null"
	nilSlice         = "[]"
	sliceStart       = "["
	sliceEnd         = "]"
)

type Number interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Any interface {
	types.ListNode | types.TreeNode
}

func Len[T Number](n T) int {
	return len(ToString(n))
}

func ToString[T Number](n T) string {
	return strconv.FormatInt(int64(n), 10)
}
