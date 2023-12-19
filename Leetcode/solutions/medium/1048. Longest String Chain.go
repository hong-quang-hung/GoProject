package medium

import (
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/longest-string-chain/
func init() {
	Solutions[1048] = func() {
		fmt.Println(`Input: words = ["a","b","ba","bca","bda","bdca"]`)
		fmt.Println(`Output:`, longestStrChain([]string{`a`, `b`, `ba`, `bca`, `bda`, `bdca`}))
		fmt.Println(`Input: words = ["xbc","pcxbcf","xb","cxbc","pcxbc"]`)
		fmt.Println(`Output:`, longestStrChain([]string{`xbc`, `pcxbcf`, `xb`, `cxbc`, `pcxbc`}))
	}
}

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	res := 0
	dp := make(map[string]int)
	for _, word := range words {
		val := 1
		for i := 0; i < len(word); i++ {
			prev := word[:i] + word[i+1:]
			if v, ok := dp[prev]; ok {
				val = max(val, v+1)
			}
		}
		res = max(res, val)
		dp[word] = val
	}
	return res
}
