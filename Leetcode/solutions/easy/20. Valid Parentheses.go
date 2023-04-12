package easy

import "fmt"

// Reference: https://leetcode.com/problems/valid-parentheses/
func Leetcode_Is_Valid() {
	fmt.Println("Input: s = '()[]{}'")
	fmt.Println("Output:", isValid("()[]{}"))
}

func isValid(s string) bool {
	var stack = []rune{}
	for _, c := range s {
		if c == '(' {
			stack = append([]rune{')'}, stack...)
		} else if c == '[' {
			stack = append([]rune{']'}, stack...)
		} else if c == '{' {
			stack = append([]rune{'}'}, stack...)
		} else {
			if len(stack) == 0 {
				return false
			}

			last := stack[0]
			stack = stack[1:]
			if last != c {
				return false
			}
		}
	}
	return len(stack) == 0
}
