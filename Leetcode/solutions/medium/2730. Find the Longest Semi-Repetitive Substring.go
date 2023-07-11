package medium

import "fmt"

// Reference: https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/
func init() {
	Solutions[2730] = func() {
		fmt.Println("Input: s = '52233'")
		fmt.Println("Output:", longestSemiRepetitiveSubstring("52233"))
		fmt.Println("Input: s = '5494'")
		fmt.Println("Output:", longestSemiRepetitiveSubstring("5494"))
		fmt.Println("Input: s = '52233256'")
		fmt.Println("Output:", longestSemiRepetitiveSubstring("52233256"))
	}
}

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	arr := []int{}
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			arr = append(arr, i)
		}
	}

	arr = append([]int{0}, arr...)
	arr = append(arr, n)

	if len(arr) == 2 {
		return n
	}

	res := 0
	for i := 2; i < len(arr); i++ {
		res = max(res, arr[i]-arr[i-2])
	}
	return res
}
