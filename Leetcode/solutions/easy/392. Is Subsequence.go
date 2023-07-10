package easy

import "fmt"

func init() {
	Solutions[392] = Leetcode_Is_Subsequence
}

// Reference: https://leetcode.com/problems/is-subsequence/
func Leetcode_Is_Subsequence() {
	fmt.Println("Input: s = 'abc', t = 'ahbgdc'")
	fmt.Println("Output:", isSubsequence("abc", "ahbgdc"))
	fmt.Println("Input: s = 'axc', t = 'ahbgdc'")
	fmt.Println("Output:", isSubsequence("axc", "ahbgdc"))
	fmt.Println("Input: s = 'acb', t = 'ahbgdc'")
	fmt.Println("Output:", isSubsequence("acb", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
