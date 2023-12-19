package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
func init() {
	Solutions[712] = func() {
		fmt.Println(`Input: s1 = "sea", s2 = "eat"`)
		fmt.Println(`Output:`, minimumDeleteSum(`sea`, `eat`))
		fmt.Println(`Input: s1 = "delete", s2 = "leet"`)
		fmt.Println(`Output:`, minimumDeleteSum(`delete`, `leet`))
	}
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	sum := 0
	for _, s := range s1 {
		sum += int(s)
	}

	for _, s := range s2 {
		sum += int(s)
	}
	return sum - 2*dp[m][n]
}
