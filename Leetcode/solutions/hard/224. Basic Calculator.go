package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/basic-calculator/
func init() {
	Solutions[224] = func() {
		fmt.Println(`Input: s = "1 + 1"`)
		fmt.Println(`Output:`, calculate(`1 + 1`))
		fmt.Println(`Input: s = " 2-1 + 2 "`)
		fmt.Println(`Output:`, calculate(` 2-1 + 2 `))
		fmt.Println(`Input: s = "(1+(4+5+2)-3)+(6+8)"`)
		fmt.Println(`Output:`, calculate(`(1+(4+5+2)-3)+(6+8)`))
	}
}

func calculate(s string) int {
	var f func(s string, i int) (int, int)
	f = func(s string, i int) (int, int) {
		res, digit, op, idx, n := 0, 0, 1, i, len(s)
		for ; idx < n && s[idx] != ')'; idx++ {
			ch := s[idx]
			switch {
			case ch == '+':
				res, digit = res+op*digit, 0
				op = 1
			case ch == '-':
				res, digit = res+op*digit, 0
				op = -1
			case ch == '(':
				digit, idx = f(s, idx+1)
			case ch >= '0' && ch <= '9':
				digit = digit*10 + int(ch-'0')
			default:
			}
		}
		return res + op*digit, idx
	}

	res, _ := f(s, 0)
	return res
}
