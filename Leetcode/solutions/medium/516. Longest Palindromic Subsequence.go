package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/longest-palindromic-subsequence/
func Leetcode_Find_Maximized_Capital() {
	fmt.Println("Input: s = 'bbbab'")
	fmt.Println("Output:", longestPalindromeSubseq("bbbab"))
	fmt.Println("Input: s = 'abcabcabcabc'")
	fmt.Println("Output:", longestPalindromeSubseq("abcabcabcabc"))
	fmt.Println("Input: s = 'aabaa'")
	fmt.Println("Output:", longestPalindromeSubseq("aabaa")) //7
}

func longestPalindromeSubseq(s string) int {
	n, maxDup := len(s), 0
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
		if maxDup < m[s[i]] {
			maxDup = m[s[i]]
		}
	}

	dp := make([][]int, n)
	for i := 1; i < n; i++ {
		dp[i] = make([]int, n)
	}
	return maxDup
}

func isPalindromeSubseq(s string, n int) bool {
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
