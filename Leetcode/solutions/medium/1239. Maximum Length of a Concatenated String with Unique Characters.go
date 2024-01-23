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
	unique := func(str string) int {
		var m [26]int
		for _, ch := range str {
			m[ch-'a'] += 1
			if m[ch-'a'] > 1 {
				return 0
			}
		}

		count := 0
		for _, val := range m {
			if val == 1 {
				count += 1
			}
		}
		return count
	}

	var f func(index int, arr []string, str string) int
	f = func(index int, arr []string, str string) int {
		if index == len(arr) {
			return unique(str)
		}
		return max(f(index+1, arr, str+arr[index]), f(index+1, arr, str))
	}
	return f(0, arr, "")
}
