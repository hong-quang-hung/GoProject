package medium

import "fmt"

func init() {
	Solutions[3] = Leetcode_Length_Of_Longest_Substring
}

// Reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func Leetcode_Length_Of_Longest_Substring() {
	fmt.Println("Input: s = 'abcabcbb'")
	fmt.Println("Output:", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println()
	fmt.Println("Input: s = 'abcdc'")
	fmt.Println("Output:", lengthOfLongestSubstring("abcdc"))
}

func lengthOfLongestSubstring(s string) int {
	chars := [128]int{}
	left, right, res := 0, 0, 0
	for right < len(s) {
		char := s[right]
		index := chars[char] - 1
		if (index != 0 || char == s[0]) && index >= left && index < right {
			left = index + 1
		}
		if res < (right - left + 1) {
			res = right - left + 1
		}
		chars[char] = right + 1
		right++
	}
	return res
}
