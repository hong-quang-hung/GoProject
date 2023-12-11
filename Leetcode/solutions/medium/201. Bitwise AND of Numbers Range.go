package medium

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/bitwise-and-of-numbers-range/
func init() {
	Solutions[201] = func() {
		fmt.Println("Input: left = 5, right = 7")
		fmt.Println("Output:", rangeBitwiseAnd(5, 7))
		fmt.Println("Input: left = 1, right = 2147483647")
		fmt.Println("Output:", rangeBitwiseAnd(1, 2147483647))
	}
}

func rangeBitwiseAnd(left int, right int) int {
	fmt.Println(strconv.FormatInt(int64(left), 2))
	fmt.Println(strconv.FormatInt(int64(right), 2))
	return 1
}
