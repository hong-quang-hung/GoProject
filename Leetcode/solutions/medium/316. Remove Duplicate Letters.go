package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/remove-duplicate-letters/
func init() {
	Solutions[316] = func() {
		fmt.Println("Input: s = 'bcabc'")
		fmt.Println("Output:", removeDuplicateLetters("bcabc"))
		fmt.Println("Input: s = 'cbacdcbc'")
		fmt.Println("Output:", removeDuplicateLetters("cbacdcbc"))
	}
}

func removeDuplicateLetters(s string) string {
	res := []rune{}
	m := [26]bool{}
	for _, ch := range s {
		i := int(ch - 'a')
		if !m[i] {
			m[i] = true
			res = append(res, ch)
		}
	}
	return string(res)
}
