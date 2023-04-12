package easy

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/word-pattern/
func Leetcode_Word_Pattern() {
	fmt.Println("Input: pattern = 'abba', s = 'dog cat cat dog'")
	fmt.Println("Output:", wordPattern("abba", "dog cat cat dog"))
}

func wordPattern(pattern string, s string) bool {
	arr := strings.Split(s, " ")
	if len(arr) != len(pattern) {
		return false
	}

	m1 := make(map[byte]string)
	m2 := make(map[string]bool)
	for i := range pattern {
		if v, c := m1[pattern[i]]; c {
			if v != arr[i] {
				return false
			}
		} else {
			if m2[arr[i]] {
				return false
			}

			m1[pattern[i]] = arr[i]
			m2[arr[i]] = true
		}
	}
	return true
}
