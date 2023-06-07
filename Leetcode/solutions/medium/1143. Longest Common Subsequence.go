package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-common-subsequence/
func Leetcode_Longest_Common_Subsequence() {
	fmt.Println("Input: text1 = 'abcde', text2 = 'ace'")
	fmt.Println("Output:", longestCommonSubsequence("abcde", "ace"))
	fmt.Println("Input: text1 = 'abc', text2 = 'abc'")
	fmt.Println("Output:", longestCommonSubsequence("abc", "abc"))
	fmt.Println("Input: text1 = 'abc', text2 = 'def'")
	fmt.Println("Output:", longestCommonSubsequence("abc", "def"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
