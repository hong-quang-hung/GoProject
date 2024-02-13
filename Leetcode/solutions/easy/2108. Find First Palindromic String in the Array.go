package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/find-first-palindromic-string-in-the-array/
func init() {
	Solutions[2108] = func() {
		fmt.Println(`Input: words = ["abc","car","ada","racecar","cool"]`)
		fmt.Println(`Output:`, firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
		fmt.Println(`Input: words = ["notapalindrome","racecar"]`)
		fmt.Println(`Output:`, firstPalindrome([]string{"notapalindrome", "racecar"}))
	}
}

func firstPalindrome(words []string) string {
	f := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for _, word := range words {
		if f(word) {
			return word
		}
	}
	return ""
}
