package hard

import "fmt"

// Reference: https://leetcode.com/problems/longest-valid-parentheses/
func init() {
	Solutions[32] = func() {
		fmt.Println("Input: s = '(()'")
		fmt.Println("Output:", longestValidParentheses("(()"))
		fmt.Println("Input: s = ')()())'")
		fmt.Println("Output:", longestValidParentheses(")()())"))
	}
}

func longestValidParentheses(s string) int {
	res := 0
	st := []int{-1}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if len(st) == 0 {
				st = []int{i}
			} else {
				res = max(res, i-st[len(st)-1])
			}
		}
	}
	return res
}
