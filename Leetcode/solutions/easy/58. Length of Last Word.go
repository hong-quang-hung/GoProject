package easy

import (
	"fmt"
	"strings"
)

func init() {
	Solutions[58] = Leetcode_Length_Of_Last_Word
}

// Reference: https://leetcode.com/problems/length-of-last-word/
func Leetcode_Length_Of_Last_Word() {
	fmt.Println("Input: s = '   fly me   to   the moon  '")
	fmt.Println("Output:", lengthOfLastWord("   fly me   to   the moon  "))
}

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	var r int = 0
	var i int = len(s) - 1
	for i >= 0 && s[i] != ' ' {
		i--
		r++
	}
	return r
}
