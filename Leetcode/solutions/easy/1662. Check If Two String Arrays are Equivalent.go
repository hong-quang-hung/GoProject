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
	i, j, n, m := 0, 0, 0, 0
	for n < len(word1) && m < len(word2) {
		if word1[n][i] != word2[m][j] {
			return false
		}

		i++
		if i == len(word1[n]) {
			i = 0
			n++
		}

		j++
		if j == len(word2[m]) {
			j = 0
			m++
		}
	}
	return n == len(word1) && m == len(word2)
}
