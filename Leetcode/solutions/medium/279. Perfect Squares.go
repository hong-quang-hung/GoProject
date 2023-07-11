package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/perfect-squares/
func init() {
	Solutions[279] = func() {
		fmt.Println("Input: n = 12")
		fmt.Println("Output:", numSquares(12))
		fmt.Println("Input: n = 13")
		fmt.Println("Output:", numSquares(13))
	}
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = -1
	}
	return numSquaresDFS(n, dp)
}

func numSquaresDFS(n int, dp []int) int {
	if dp[n] != -1 {
		return dp[n]
	}

	x := int(math.Sqrt(float64(n)))
	count := math.MaxInt
	for i := 1; i <= x; i++ {
		count = min(count, 1+numSquaresDFS(n-i*i, dp))
	}

	dp[n] = count
	return count
}
