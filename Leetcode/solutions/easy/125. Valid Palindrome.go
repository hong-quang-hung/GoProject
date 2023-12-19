package easy

import (
	"fmt"
	"strings"
	"unicode"
)

// Reference: https://leetcode.com/problems/valid-palindrome/
func init() {
	Solutions[125] = func() {
		fmt.Println(`Input: s = "A man, a plan, a canal: Panama"`)
		fmt.Println(`Output:`, isPalindromeII(`A man, a plan, a canal: Panama`))
		fmt.Println(`Input: s = "race a car"`)
		fmt.Println(`Output:`, isPalindromeII(`race a car`))
	}
}

func isPalindromeII(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}

		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
