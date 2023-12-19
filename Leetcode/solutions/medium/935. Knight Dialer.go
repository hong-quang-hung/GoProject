package medium

import "fmt"

// Reference: https://leetcode.com/problems/bus-routes/
func init() {
	Solutions[935] = func() {
		fmt.Println(`Input: n = 1`)
		fmt.Println(`Output:`, knightDialer(1))
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, knightDialer(2))
		fmt.Println(`Input: n = 3131`)
		fmt.Println(`Output:`, knightDialer(3131))
	}
}

func knightDialer(n int) int {
	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		dp[i] = 1
	}

	if n > 1 {
		dp[5] = 0
	}

	for i := 1; i < n; i++ {
		prev := make([]int, 10)
		copy(prev, dp)
		dp[0] = (prev[4] + prev[6]) % mod
		dp[1] = (prev[6] + prev[8]) % mod
		dp[2] = (prev[7] + prev[9]) % mod
		dp[3] = (prev[4] + prev[8]) % mod
		dp[4] = (prev[0] + prev[3] + prev[9]) % mod
		dp[6] = (prev[0] + prev[1] + prev[7]) % mod
		dp[7] = (prev[2] + prev[6]) % mod
		dp[8] = (prev[1] + prev[3]) % mod
		dp[9] = (prev[2] + prev[4]) % mod
	}

	res := 0
	for _, num := range dp {
		res = (res + num) % mod
	}
	return res
}
