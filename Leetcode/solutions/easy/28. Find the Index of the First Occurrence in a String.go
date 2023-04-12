package easy

import "fmt"

// Reference: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func Leetcode_StrStr() {
	fmt.Println("Input: haystack = \"sadbutsad\", needle = \"sad\"")
	fmt.Println("Output:", strStr("sadbutsad", "sad"))
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			var j int = 1
			for ; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}
