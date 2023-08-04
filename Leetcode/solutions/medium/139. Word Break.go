package medium

import "fmt"

// Reference: https://leetcode.com/problems/word-break/
func init() {
	Solutions[139] = func() {
		fmt.Println("Input: s = 'leetcode', wordDict = ['leet','code']")
		fmt.Println("Output:", wordBreak("leetcode", []string{"leet", "code"}))
		fmt.Println("Input: s = 'applepenapple', wordDict = ['apple','pen']")
		fmt.Println("Output:", wordBreak("applepenapple", []string{"apple", "pen"}))
		fmt.Println("Input: s = 'catsandog', wordDict = ['cats','dog','sand','and','cat']")
		fmt.Println("Output:", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	}
}

func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	for _, word := range wordDict {
		m[word] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return wordBreakSolved(s, dp, m, 0, n)
}

func wordBreakSolved(s string, dp []int, m map[string]bool, idx int, n int) bool {
	if idx == n {
		return true
	}

	if dp[idx] != -1 {
		return dp[idx] == 1
	}

	for i := idx + 1; i <= n; i++ {
		if _, ok := m[s[idx:i]]; ok {
			if wordBreakSolved(s, dp, m, i, n) {
				dp[idx] = 1
				return true
			} else {
				dp[idx] = 0
			}
		} else {
			dp[idx] = 0
		}
	}
	return false
}
