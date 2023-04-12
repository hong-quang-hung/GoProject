package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-palindromic-substring/
func Leetcode_Longest_Palindrome() {
	fmt.Println("Input: s = 'babad'")
	fmt.Println("Output:", longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	var start int = 0
	var end int = 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		var len int
		if len1 > len2 {
			len = len1
		} else {
			len = len2
		}
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	l, r := left, right
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
