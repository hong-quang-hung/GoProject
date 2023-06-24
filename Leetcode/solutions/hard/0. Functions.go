package hard

import (
	"fmt"

	"leetcode.com/Leetcode/utils"
)

func min(a, b int) int {
	return utils.MinInt(a, b)
}

func max(a, b int) int {
	return utils.MaxInt(a, b)
}

func gcd(a, b int) int {
	return utils.GcdInt(a, b)
}

func abs(x int) int {
	return utils.AbsInt(x)
}

func min64(a, b int64) int64 {
	return utils.MinInt64(a, b)
}

func Leetcode_SQL() {
	fmt.Printf("This is SQL solution!\n")
}

func Leetcode_Javascript() {
	fmt.Printf("This is Javascript solution!\n")
}
