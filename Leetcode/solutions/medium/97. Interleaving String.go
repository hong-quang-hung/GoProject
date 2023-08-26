package medium

import "fmt"

// Reference: https://leetcode.com/problems/interleaving-string/
func init() {
	Solutions[97] = func() {
		fmt.Println("Input: s1 = 'aabcc', s2 = 'dbbca', s3 = 'aadbbcbcac'")
		fmt.Println("Output:", isInterleave("aabcc", "dbbca", "aadbbcbcac"))
		fmt.Println("Input: s1 = 'aabcc', s2 = 'dbbca', s3 = 'aadbbbaccc'")
		fmt.Println("Output:", isInterleave("aabcc", "dbbca", "aadbbbaccc"))
		fmt.Println("Input: s1 = '', s2 = '', s3 = ''")
		fmt.Println("Output:", isInterleave("", "", ""))
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	return false
}
