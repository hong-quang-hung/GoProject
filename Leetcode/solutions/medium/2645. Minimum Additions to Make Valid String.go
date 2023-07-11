package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/minimum-additions-to-make-valid-string/
func init() {
	Solutions[2645] = func() {
		fmt.Println("Input: word = 'aaaaab'")
		fmt.Println("Output:", addMinimum("aaaaab"))
		fmt.Println("Input: word = 'aaabcb'")
		fmt.Println("Output:", addMinimum("aaabcb"))
	}
}

func addMinimum(word string) int {
	word = strings.ReplaceAll(word, "abc", "***")
	res := 0
	pattern := []string{"ab", "bc", "ac", "bc"}
	for _, p := range pattern {
		temp := word
		temp = strings.ReplaceAll(temp, p, "")
		word = strings.ReplaceAll(word, p, "**")
		res += (len(word) - len(temp)) / 2
	}

	word = strings.ReplaceAll(word, "*", "")
	res += 2 * len(word)
	return res
}
