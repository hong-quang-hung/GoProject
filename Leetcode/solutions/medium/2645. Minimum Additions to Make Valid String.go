package medium

import (
	"fmt"
	"strings"
)

func init() {
	Solutions[2645] = Leetcode_Add_Minimum
}

// Reference: https://leetcode.com/problems/minimum-additions-to-make-valid-string/
func Leetcode_Add_Minimum() {
	fmt.Println("Input: word = 'aaaaab'")
	fmt.Println("Output:", addMinimum("aaaaab"))
	fmt.Println("Input: word = 'aaa'")
	fmt.Println("Output:", addMinimum("aaa"))
	fmt.Println("Input: word = 'aaabcb'")
	fmt.Println("Output:", addMinimum("aaabcb"))
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
