package medium

import "fmt"

// Reference: https://leetcode.com/problems/removing-stars-from-a-string/
func Leetcode_Remove_Stars() {
	fmt.Println("Input: s = 'leet**cod*e'")
	fmt.Println("Output:", removeStars("leet**cod*e"))
	fmt.Println("Input: s = 'erase*****'")
	fmt.Println("Output:", removeStars("erase*****"))
}

func removeStars(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
