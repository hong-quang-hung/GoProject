package medium

import "fmt"

// Reference: https://leetcode.com/problems/decode-ways/
func init() {
	Solutions[91] = func() {
		fmt.Println(`Input: s = "12"`)
		fmt.Println(`Output:`, numDecodings("12"))
		fmt.Println(`Input: s = "226"`)
		fmt.Println(`Output:`, numDecodings("226"))
		fmt.Println(`Input: s = "06"`)
		fmt.Println(`Output:`, numDecodings("06"))
	}
}

func numDecodings(s string) int {
	// dp := make(map[string]int)
	res := 0
	return res
}
