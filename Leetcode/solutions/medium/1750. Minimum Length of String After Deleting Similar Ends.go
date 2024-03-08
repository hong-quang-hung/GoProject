package medium

import "fmt"

// Reference: https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/
func init() {
	Solutions[1750] = func() {
		fmt.Println(`Input: s = "ca"`)
		fmt.Println(`Output:`, minimumLength("ca"))
		fmt.Println(`Input: s = "cabaabac"`)
		fmt.Println(`Output:`, minimumLength("cabaabac"))
		fmt.Println(`Input: s = "aabccabba"`)
		fmt.Println(`Output:`, minimumLength("aabccabba"))
	}
}

func minimumLength(s string) int {
	l, r := 0, len(s)-1
	for (l < r) && s[l] == s[r] {
		ch := s[l]
		for ; l <= r && s[l] == ch; l++ {
		}
		for ; l <= r && s[r] == ch; r-- {
		}
	}
	return r - l + 1
}
