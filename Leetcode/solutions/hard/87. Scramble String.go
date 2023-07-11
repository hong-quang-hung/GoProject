package hard

import "fmt"

// Reference: https://leetcode.com/problems/scramble-string/
func init() {
	Solutions[87] = func() {
		fmt.Println("Input: s1 = 'great', s2 = 'rgeat'")
		fmt.Println("Output:", isScramble("great", "rgeat"))
		fmt.Println("Input: s1 = 'abcde', s2 = 'caebd'")
		fmt.Println("Output:", isScramble("abcde", "caebd"))
	}
}

func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	dp := make([][][]bool, n+1)
	dp[1] = make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[1][i] = make([]bool, n)
		for j := 0; j < n; j++ {
			dp[1][i][j] = s1[i] == s2[j]
		}
	}

	for l := 2; l <= n; l++ {
		if dp[l] == nil {
			dp[l] = make([][]bool, n)
		}

		for i := 0; i < n+1-l; i++ {
			if dp[l][i] == nil {
				dp[l][i] = make([]bool, n)
			}

			for j := 0; j < n+1-l; j++ {
				for k := 1; k < l; k++ {
					dp1 := dp[k][i]
					dp2 := dp[l-k][i+k]

					dp[l][i][j] = dp[l][i][j] || (dp1[j] && dp2[j+k])
					dp[l][i][j] = dp[l][i][j] || (dp1[j+l-k] && dp2[j])
				}
			}
		}
	}
	return dp[n][0][0]
}
