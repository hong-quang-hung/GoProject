package easy

import "fmt"

// Reference: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
func init() {
	Solutions[1662] = func() {
		fmt.Println("Input: word1 = {\"ab\", \"c\"}, word2 = {\"a\", \"bc\"}")
		fmt.Println("Output:", arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))
		fmt.Println("Input: word1 = {\"a\", \"cb\"}, word2 = {\"ab\", \"c\"}")
		fmt.Println("Output:", arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))
		fmt.Println("Input: word1  = {\"abc\", \"d\", \"defg\"}, word2 = {\"abcddefg\"}")
		fmt.Println("Output:", arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"}))
	}
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return true
}
