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
	n := len(s)
	ms, mt := make([]int, 26), make([]int, 26)
	for i := 0; i < n; i++ {
		ms[int(s[i]-'a')]++
		mt[int(t[i]-'a')]++
	}

	fmt.Println(ms)
	fmt.Println(mt)

	res := 0
	for i := 0; i < 26; i++ {

	}
	return res
}
