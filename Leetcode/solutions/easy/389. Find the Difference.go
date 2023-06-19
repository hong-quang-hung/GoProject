package easy

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/find-the-difference/
func Leetcode_Find_The_Difference() {
	fmt.Println("Input: s = 'abcd', t = 'abcde'")
	fmt.Println("Output:", findTheDifference("abcd", "abcde"))
	fmt.Println("Input: s = '', t = 'y'")
	fmt.Println("Output:", findTheDifference("", "y"))
}

func findTheDifference(s string, t string) byte {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	for i := range t {
		if m[t[i]] == 0 {
			return t[i]
		}
		m[t[i]]--
	}
	return ' '
}
