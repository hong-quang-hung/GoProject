package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/
func init() {
	Solutions[459] = func() {
		fmt.Println("Input: s = 'abab'")
		fmt.Println("Output:", repeatedSubstringPattern("abab"))
		fmt.Println("Input: s = 'aba'")
		fmt.Println("Output:", repeatedSubstringPattern("aba"))
		fmt.Println("Input: s = 'abcabcabcabc'")
		fmt.Println("Output:", repeatedSubstringPattern("abcabcabcabc"))
	}
}

func repeatedSubstringPattern(s string) bool {
	t := s + s
	return strings.Contains(t[1:len(t)-1], s)
}
