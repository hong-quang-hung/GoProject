package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/basic-calculator/
func init() {
	Solutions[224] = func() {
		fmt.Println("Input: s = '1 + 1'")
		fmt.Println("Output:", calculate("1 + 1"))
		fmt.Println("Input: s = ' 2-1 + 2 '")
		fmt.Println("Output:", calculate(" 2-1 + 2 "))
		fmt.Println("Input: s = '(1+(4+5+2)-3)+(6+8)'")
		fmt.Println("Output:", calculate("(1+(4+5+2)-3)+(6+8)"))
	}
}

func calculate(s string) int {
	res := 0
	return res
}
