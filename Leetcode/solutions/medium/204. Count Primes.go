package medium

import "fmt"

// Reference: https://leetcode.com/problems/count-primes/
func init() {
	Solutions[204] = func() {
		fmt.Println(`Input: n = 10`)
		fmt.Println(`Output:`, countPrimes(10))
		fmt.Println(`Input: n = 994794`)
		fmt.Println(`Output:`, countPrimes(994794))
	}
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	prime := make([]bool, n)
	for i := 0; i < n; i++ {
		prime[i] = true
	}

	res := 1
	for i := 3; i < n; i += 2 {
		if prime[i] {
			res++
		}

		for j := 2 * i; j < n; j = j + i {
			prime[j] = false
		}
	}
	return res
}
