package medium

import (
	"fmt"
	"strings"
)

// Reference: https://leetcode.com/problems/reverse-words-in-a-string/
func Leetcode_Reverse_Words() {
	fmt.Println("Input: s = 'the sky is blue'")
	fmt.Println("Output:", reverseWords("the sky is blue"))
	fmt.Println("Input: s = '  hello world  '")
	fmt.Println("Output:", reverseWords("  hello world  "))
	fmt.Println("Input: s = 'a good   example'")
	fmt.Println("Output:", reverseWords("a good   example"))
}

func reverseWords(s string) string {
	res := strings.Fields(s)
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return strings.Join(res, " ")
}
