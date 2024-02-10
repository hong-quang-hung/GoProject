package medium

import "fmt"

// Reference: https://leetcode.com/problems/palindromic-substrings/
func init() {
	Solutions[647] = func() {
		fmt.Println(`Input: s = "abc"`)
		fmt.Println(`Output:`, countSubstrings("abc"))
		fmt.Println(`Input: s = "aaa"`)
		fmt.Println(`Output:`, countSubstrings("aaa"))
	}
}

func countSubstrings(s string) int {
	f := func(l, r int) int {
		count := 0
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
		return count
	}

	res := 0
	for i := range s {
		res += f(i, i) + f(i, i+1)
	}
	return res
}
