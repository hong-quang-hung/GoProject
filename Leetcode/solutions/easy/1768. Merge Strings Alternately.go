package easy

import "fmt"

// Reference: https://leetcode.com/problems/merge-strings-alternately/
func init() {
	Solutions[1768] = func() {
		fmt.Println("Input: word1 = 'abc', word2 = 'pqr'")
		fmt.Println("Output:", mergeAlternately("abc", "pqr"))
		fmt.Println("Input: word1 = 'ab', word2 = 'pqrs'")
		fmt.Println("Output:", mergeAlternately("ab", "pqrs"))
		fmt.Println("Input: word1 = 'abcd', word2 = 'pq'")
		fmt.Println("Output:", mergeAlternately("abcd", "pq"))
	}
}

func mergeAlternately(word1 string, word2 string) string {
	n, m, sub := len(word1), len(word2), ""
	l := n
	if n > m {
		sub = word1[m:]
		l = m
	} else if n < m {
		sub = word2[n:]
	}

	str := []byte{}
	for i := 0; i < l; i++ {
		str = append(str, ([]byte{word1[i], word2[i]})...)
	}
	return string(str) + sub
}
