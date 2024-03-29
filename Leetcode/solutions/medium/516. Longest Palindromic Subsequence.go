package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-palindromic-subsequence/
func init() {
	Solutions[516] = func() {
		fmt.Println(`Input: s = "bbbab"`)
		fmt.Println(`Output:`, longestPalindromeSubseq(`bbbab`))
		fmt.Println(`Input: s = "abcabcabcabc"`)
		fmt.Println(`Output:`, longestPalindromeSubseq(`abcabcabcabc`))
		fmt.Println(`Input: s = "aabaa"`)
		fmt.Println(`Output:`, longestPalindromeSubseq(`aabaa`))
	}
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp, dpPrev := make([]int, n), make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[j] = dpPrev[j-1] + 2
			} else {
				dp[j] = max(dpPrev[j], dp[j-1])
			}
		}
		copy(dpPrev, dp)
	}
	return dp[n-1]
}
