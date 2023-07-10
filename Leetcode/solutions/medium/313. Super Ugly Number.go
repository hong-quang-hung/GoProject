package medium

import (
	"fmt"
	"math"
)

func init() {
	Solutions[313] = Leetcode_Nth_Super_Ugly_Number
}

// Reference: https://leetcode.com/problems/super-ugly-number/
func Leetcode_Nth_Super_Ugly_Number() {
	fmt.Println("Input: n = 12, primes = [2,7,13,19]")
	fmt.Println("Output:", nthSuperUglyNumber(12, []int{2, 7, 13, 19}))
	fmt.Println("Input: n = 1, primes = [2,3,5]")
	fmt.Println("Output:", nthSuperUglyNumber(1, []int{2, 3, 5}))
}

func nthSuperUglyNumber(n int, primes []int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	p := len(primes)
	dp[1] = 1
	primeVal := make([]int, p)
	for i := 0; i < p; i++ {
		primeVal[i] = 1
	}

	for i := 2; i <= n; i++ {
		minP := math.MaxInt
		for j := 0; j < p; j++ {
			minP = min(minP, dp[primeVal[j]]*primes[j])
		}

		dp[i] = minP

		for j := 0; j < p; j++ {
			if dp[i] == dp[primeVal[j]]*primes[j] {
				primeVal[j]++
			}
		}
	}
	return dp[n]
}
