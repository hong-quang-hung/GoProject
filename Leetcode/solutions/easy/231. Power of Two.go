package easy

import "fmt"

// Reference: https://leetcode.com/problems/power-of-two/
func Leetcode_Is_Power_Of_Two() {
	fmt.Println("Input: n = 16")
	fmt.Println("Output:", isPowerOfTwo(16))
	fmt.Println("Input: n = 3")
	fmt.Println("Output:", isPowerOfTwo(3))
}

func isPowerOfTwo(n int) bool {
	if n < 2 {
		return n == 1
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}

		n /= 2
	}
	return true
}
