package medium

import "fmt"

// Reference: https://leetcode.com/problems/extra-characters-in-a-string/
func init() {
	Solutions[2707] = func() {
		fmt.Println("Input: s = 'leetscode', dictionary = ['leet','code','leetcode']")
		fmt.Println("Output:", minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
		fmt.Println("Input: s = 'sayhelloworld', dictionary = ['hello','world']")
		fmt.Println("Output:", minExtraChar("sayhelloworld", []string{"hello", "world"}))
	}
}

func minExtraChar(s string, dictionary []string) int {
	dict := make(map[string]bool)
	for _, d := range dictionary {
		dict[d] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			subStr := s[j:i]
			if dict[subStr] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}
