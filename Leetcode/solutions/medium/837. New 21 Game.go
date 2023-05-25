package medium

import "fmt"

// Reference: https://leetcode.com/problems/new-21-game/
func Leetcode_New_21_Game() {
	fmt.Println("Input: n = 10, k = 1, maxPts = 10")
	fmt.Println("Output:", new21Game(10, 1, 10))
	fmt.Println("Input: n = 6, k = 1, maxPts = 10")
	fmt.Println("Output:", new21Game(6, 1, 10))
	fmt.Println("Input: n = 21, k = 17, maxPts = 10")
	fmt.Println("Output:", new21Game(21, 17, 10))
}

func new21Game(n int, k int, maxPts int) float64 {
	dp, s := make([]float64, n+1), float64(1)
	dp[0] = 1
	if k <= 0 {
		s = 0
	}

	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			s -= dp[i-maxPts]
		}
	}

	ans := float64(0)
	for i := k; i <= n; i++ {
		ans += dp[i]
	}
	return ans
}
