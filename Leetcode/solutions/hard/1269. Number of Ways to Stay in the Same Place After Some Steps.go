package hard

import "fmt"

// Reference: https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
func init() {
	Solutions[1269] = func() {
		fmt.Println(`Input: steps = 3, arrLen = 2`)
		fmt.Println(`Output:`, numWays(3, 2))
		fmt.Println(`Input: steps = 2, arrLen = 4`)
		fmt.Println(`Output:`, numWays(2, 4))
		fmt.Println(`Input: steps = 4, arrLen = 2`)
		fmt.Println(`Output:`, numWays(4, 2))
	}
}

func numWays(steps int, arrLen int) int {
	if arrLen > steps {
		arrLen = steps
	}

	dp := make([]int, arrLen)
	prevDp := make([]int, arrLen)
	prevDp[0] = 1
	for i := 1; i <= steps; i++ {
		dp = make([]int, arrLen)
		for j := arrLen - 1; j >= 0; j-- {
			res := prevDp[j]
			if j > 0 {
				res = (res + prevDp[j-1]) % mod
			}
			if j < arrLen-1 {
				res = (res + prevDp[j+1]) % mod
			}
			dp[j] = res
		}
		prevDp = dp
	}
	return dp[0]
}
