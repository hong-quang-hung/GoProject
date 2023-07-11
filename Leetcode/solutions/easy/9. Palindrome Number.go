package easy

import "fmt"

// Reference: https://leetcode.com/problems/palindrome-number/
func init() {
	Solutions[9] = func() {
		fmt.Println("Input: x = 0")
		fmt.Println("Output:", isPalindrome(0))
	}
}

func isPalindrome(x int) bool {
	r := 0
	t := x
	for x > 0 {
		r = (r * 10) + (x % 10)
		x = x / 10

	}
	return t == r
}
