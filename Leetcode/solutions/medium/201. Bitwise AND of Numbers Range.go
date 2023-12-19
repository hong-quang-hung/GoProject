package medium

import (
	"fmt"
	"strconv"
)

// Reference: https://leetcode.com/problems/bitwise-and-of-numbers-range/
func init() {
	Solutions[201] = func() {
		fmt.Println(`Input: left = 5, right = 7`)
		fmt.Println(`Output:`, rangeBitwiseAnd(5, 7))
		fmt.Println(`Input: left = 1, right = 2147483647`)
		fmt.Println(`Output:`, rangeBitwiseAnd(1, 2147483647))
		fmt.Println(`Input: left = 5, right = 5`)
		fmt.Println(`Output:`, rangeBitwiseAnd(5, 5))
	}
}

func rangeBitwiseAnd(left int, right int) int {
	leftStr := strconv.FormatInt(int64(left), 2)
	rightStr := strconv.FormatInt(int64(right), 2)

	if len(leftStr) != len(rightStr) {
		return 0
	}

	res := 0
	for i := 0; i < len(leftStr); i++ {
		if leftStr[i] == '1' && rightStr[i] == '1' {
			res += 1 << (len(leftStr) - i - 1)
		} else if leftStr[i] != rightStr[i] {
			break
		}
	}
	return res
}
