package easy

import "fmt"

// Reference: https://leetcode.com/problems/is-subsequence/
func init() {
	Solutions[392] = func() {
		fmt.Println(`Input: s = "abc", t = "ahbgdc"`)
		fmt.Println(`Output:`, isSubsequence(`abc`, `ahbgdc`))
		fmt.Println(`Input: s = "axc", t = "ahbgdc"`)
		fmt.Println(`Output:`, isSubsequence(`axc`, `ahbgdc`))
	}
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
