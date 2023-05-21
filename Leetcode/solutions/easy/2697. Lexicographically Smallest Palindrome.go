package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/lexicographically-smallest-palindrome/
func Leetcode_Make_Smallest_Palindrome() {
	fmt.Println("Input: s = 'egcfe'")
	fmt.Println("Output:", makeSmallestPalindrome("egcfe"))
	fmt.Println("Input: s = 'abcd'")
	fmt.Println("Output:", makeSmallestPalindrome("abcd"))
	fmt.Println("Input: s = 'seven'")
	fmt.Println("Output:", makeSmallestPalindrome("seven"))
}

func makeSmallestPalindrome(s string) string {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] > s[n-1-i] {
			s = replaceAtIndex(s, s[n-1-i], i)
		} else {
			s = replaceAtIndex(s, s[i], n-1-i)
		}
	}
	return s
}

func replaceAtIndex(s string, r byte, i int) string {
	out := []byte(s)
	out[i] = r
	return string(out)
}
