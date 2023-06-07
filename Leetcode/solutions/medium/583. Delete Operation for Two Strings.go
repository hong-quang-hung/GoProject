package medium

import "fmt"

// Reference: https://leetcode.com/problems/delete-operation-for-two-strings/
func Leetcode_Min_Distance() {
	fmt.Println("Input: word1 = 'sea', word2 = 'eat'")
	fmt.Println("Output:", minDistance("sea", "eat"))
	fmt.Println("Input: word1 = 'leetcode', word2 = 'etco'")
	fmt.Println("Output:", minDistance("leetcode", "etco"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return m + n - 2*dp[m][n]
}
