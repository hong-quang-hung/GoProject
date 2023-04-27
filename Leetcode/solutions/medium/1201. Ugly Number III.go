package medium

import "fmt"

// Reference: https://leetcode.com/problems/ugly-number-iii/
func Leetcode_Is_Ugly_III() {
	fmt.Println("Input: n = 3, a = 2, b = 3, c = 5")
	fmt.Println("Output:", nthUglyNumber_iii(3, 2, 3, 5))
	fmt.Println("Input: n = 4, a = 2, b = 3, c = 4")
	fmt.Println("Output:", nthUglyNumber_iii(4, 2, 3, 4))
	fmt.Println("Input: n = 5, a = 2, b = 11, c = 13")
	fmt.Println("Output:", nthUglyNumber_iii(5, 2, 11, 13))
}

func nthUglyNumber_iii(n int, a int, b int, c int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	x, y, z := 1, 1, 1
	for i := 1; i <= n; i++ {
		dp[i] = min(min(dp[x-1]*a, dp[y-1]*b), dp[z-1]*c)
		// fmt.Println(dp[i], i, x, y, z, dp)
		if dp[i] == dp[x-1]*a {
			x++
		}
		if dp[i] == dp[y-1]*b {
			y++
		}
		if dp[i] == dp[z-1]*c {
			z++
		}
	}
	fmt.Println(dp)
	return dp[n]
}
