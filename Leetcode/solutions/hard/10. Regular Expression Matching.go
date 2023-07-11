package hard

import "fmt"

// Reference: https://leetcode.com/problems/regular-expression-matching/
func init() {
	Solutions[10] = func() {
		fmt.Println("Input: s = 'aa', p = 'a'")
		fmt.Println("Output:", isMatch("aa", "a"))
		fmt.Println("Input: s = 'aa', p = 'a*'")
		fmt.Println("Output:", isMatch("aa", "a*"))
		fmt.Println("Input: s = 'ab', p = '.*'")
		fmt.Println("Output:", isMatch("ab", ".*"))
	}
}

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[n][m] = true
	for i := n; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			match := (i < n && (p[j] == s[i] || p[j] == '.'))
			if j+1 < m && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || match && dp[i+1][j]
			} else {
				dp[i][j] = match && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}
