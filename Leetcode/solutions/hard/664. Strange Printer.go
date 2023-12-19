package hard

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/strange-printer/
func init() {
	Solutions[664] = func() {
		fmt.Println(`Input: s = "aaabbb"`)
		fmt.Println(`Output:`, strangePrinter(`aaabbb`))
		fmt.Println(`Input: s = "aba"`)
		fmt.Println(`Output:`, strangePrinter(`aba`))
	}
}

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = math.MaxInt64
				for k := i; k < j; k++ {
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}
