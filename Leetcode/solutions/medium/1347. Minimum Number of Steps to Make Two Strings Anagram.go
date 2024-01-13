package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
func init() {
	Solutions[1347] = func() {
		fmt.Println(`Input: s = "bab", t = "aba"`)
		fmt.Println(`Output:`, minSteps("bab", "aba"))
		fmt.Println(`Input: s = "leetcode", t = "practice"`)
		fmt.Println(`Output:`, minSteps("leetcode", "practice"))
		fmt.Println(`Input: s = "anagram", t = "mangaar"`)
		fmt.Println(`Output:`, minSteps("anagram", "mangaar"))
	}
}

func minSteps(s string, t string) int {
	m := make([]int, 26)
	for i := range s {
		m[s[i]-'a']++
	}

	res := 0
	for i := range t {
		if m[t[i]-'a'] == 0 {
			res++
			continue
		}
		m[t[i]-'a']--
	}
	return res
}
