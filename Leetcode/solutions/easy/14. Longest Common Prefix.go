package easy

import (
	"fmt"
	"strings"
)

func init() {
	Solutions[14] = Leetcode_Longest_Common_Prefix
}

// Reference: https://leetcode.com/problems/longest-common-prefix/
func Leetcode_Longest_Common_Prefix() {
	fmt.Println("Input: strs = ['flower','flow','flight']")
	fmt.Println("Output:", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("Input: strs = ['dog','racecar','car']")
	fmt.Println("Output:", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	prefix := strs[0]
	for _, s := range strs {
		for strings.Index(s, prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}
