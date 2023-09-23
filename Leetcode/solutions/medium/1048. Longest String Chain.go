package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-string-chain/
func init() {
	Solutions[1048] = func() {
		fmt.Println("Input: words = ['a','b','ba','bca','bda','bdca']")
		fmt.Println("Output:", longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
		fmt.Println("Input: words = ['xbc','pcxbcf','xb','cxbc','pcxbc']")
		fmt.Println("Output:", longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
		fmt.Println("Input: words = ['abcd','dbqca']")
		fmt.Println("Output:", longestStrChain([]string{"abcd", "dbqca"}))
	}
}

func longestStrChain(words []string) int {
	return 0
}
