package medium

import "fmt"

// Reference: https://leetcode.com/problems/reorganize-string/
func init() {
	Solutions[767] = func() {
		fmt.Println("Input: s = 'aab'")
		fmt.Println("Output:", reorganizeString("aab"))
		fmt.Println("Input: s = 'aaab'")
		fmt.Println("Output:", reorganizeString("aaab"))
	}
}

func reorganizeString(s string) string {
	return ""
}
