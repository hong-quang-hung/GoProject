package medium

import "fmt"

// Reference: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func init() {
	Solutions[3] = func() {
		fmt.Println(`Input: s = "abcabcbb"`)
		fmt.Println(`Output:`, lengthOfLongestSubstring(`abcabcbb`))
		fmt.Println(`Input: s = "abcdc"`)
		fmt.Println(`Output:`, lengthOfLongestSubstring(`abcdc`))
	}
}

func lengthOfLongestSubstring(s string) int {
	chars, left, right, res := [128]int{}, 0, 0, 0
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
