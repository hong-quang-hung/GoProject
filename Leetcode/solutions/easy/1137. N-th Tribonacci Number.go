package easy

import "fmt"

func init() {
	Solutions[1137] = Leetcode_Tribonacci
}

// Reference: https://leetcode.com/problems/n-th-tribonacci-number/
func Leetcode_Tribonacci() {
	fmt.Println("Input: n = 4")
	fmt.Println("Output:", tribonacci(4))
	fmt.Println("Input: n = 25")
	fmt.Println("Output:", tribonacci(25))
	fmt.Println("Input: n = 2")
	fmt.Println("Output:", tribonacci(2))
}

func tribonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	if n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
