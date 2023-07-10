package easy

import (
	"fmt"
	"strings"
)

func init() {
	Solutions[2609] = Leetcode_Find_The_Longest_Balanced_Substring
}

// Reference: https://leetcode.com/problems/find-the-longest-balanced-substring-of-a-binary-string/
func Leetcode_Find_The_Longest_Balanced_Substring() {
	fmt.Println("Input: s = '01011'")
	fmt.Println("Output:", findTheLongestBalancedSubstring("01011"))
}

func findTheLongestBalancedSubstring(s string) int {
	res := 0
	temp := "01"
	for len(temp) <= len(s) {
		if strings.Contains(s, temp) {
			res = len(temp)
		}
		temp = "0" + temp + "1"
	}
	return res
}
