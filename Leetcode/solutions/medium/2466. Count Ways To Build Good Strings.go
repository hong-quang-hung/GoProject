package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/count-ways-to-build-good-strings/
func Leetcode_Count_Good_Strings() {
	fmt.Println("Input: low = 3, high = 3, zero = 1, one = 1")
	fmt.Println("Output:", countGoodStrings(3, 3, 1, 1))
	fmt.Println("Input: low = 2, high = 3, zero = 1, one = 2")
	fmt.Println("Output:", countGoodStrings(2, 3, 1, 2))
}

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	for i := 1; i <= high; i++ {
		dp[i] = -1
	}

	dp[0] = 1
	res := 0
	for i := low; i <= high; i++ {
		res += countGoodStringsSolved(zero, one, dp, i)
		res %= mod
	}
	return res
}

func countGoodStringsSolved(zero int, one int, dp []int, i int) int {
	if dp[i] != -1 {
		return dp[i]
	}

	count := 0
	if i >= one {
		count += countGoodStringsSolved(zero, one, dp, i-one)
	}

	if i >= zero {
		count += countGoodStringsSolved(zero, one, dp, i-zero)
	}

	dp[i] = count % mod
	return dp[i]
}
