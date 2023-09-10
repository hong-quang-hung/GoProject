package medium

import "fmt"

// Reference: https://leetcode.com/problems/interleaving-string/
func init() {
	Solutions[97] = func() {
		fmt.Println("Input: s1 = 'aabcc', s2 = 'dbbca', s3 = 'aadbbcbcac'")
		fmt.Println("Output:", isInterleave("aabcc", "dbbca", "aadbbcbcac"))
		fmt.Println("Input: s1 = 'aabcc', s2 = 'dbbca', s3 = 'aadbbbaccc'")
		fmt.Println("Output:", isInterleave("aabcc", "dbbca", "aadbbbaccc"))
		fmt.Println("Input: s1 = '', s2 = '', s3 = ''")
		fmt.Println("Output:", isInterleave("", "", ""))
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i += 1 {
		for j := 0; j <= len(s2); j += 1 {
			k := i + j - 1
			dp[j] = (i == 0 && j == 0) || i > 0 && s3[k] == s1[i-1] && dp[j] || j > 0 && s3[k] == s2[j-1] && dp[j-1]
		}
	}
	return dp[len(s2)]
}
