package easy

import "fmt"

// Reference: https://leetcode.com/problems/fibonacci-number/
func Leetcode_Fibonacci() {
	fmt.Println("Input: n = 2")
	fmt.Println("Output:", fib(2))
	fmt.Println("Input: n = 4")
	fmt.Println("Output:", fib(4))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = -1
	}
	return fib_solved(dp, n)
}

func fib_solved(dp []int, i int) int {
	if dp[i] != -1 {
		return dp[i]
	}

	dp[i] = fib_solved(dp, i-1) + fib_solved(dp, i-2)
	return dp[i]
}
