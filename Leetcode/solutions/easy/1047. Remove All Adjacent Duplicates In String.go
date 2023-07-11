package easy

import "fmt"

// Reference: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func init() {
	Solutions[1047] = func() {
		fmt.Println("Input: s = 'azxxzy'")
		fmt.Println("Output:", removeAdjacentDuplicates("azxxzy"))
	}
}

func removeAdjacentDuplicates(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
