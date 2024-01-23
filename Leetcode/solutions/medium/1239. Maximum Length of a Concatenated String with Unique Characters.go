package medium

import "fmt"

// Reference: https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
func init() {
	Solutions[1239] = func() {
		fmt.Println(`Input: arr = ["un","iq","ue"]`)
		fmt.Println(`Output:`, maxLength([]string{"un", "iq", "ue"}))
		fmt.Println(`Input: arr = ["cha","r","act","ers"]`)
		fmt.Println(`Output:`, maxLength([]string{"cha", "r", "act", "ers"}))
		fmt.Println(`Input: arr = ["abcdefghijklmnopqrstuvwxyz"]`)
		fmt.Println(`Output:`, maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}))
	}
}

func maxLength(arr []string) int {
	return 0
}
