package easy

import "fmt"

func init() {
	Solutions[326] = Leetcode_Is_Power_Of_Three
}

// Reference: https://leetcode.com/problems/power-of-three/
func Leetcode_Is_Power_Of_Three() {
	fmt.Println("Input: n = 27")
	fmt.Println("Output:", isPowerOfThree(27))
	fmt.Println("Input: n = -1")
	fmt.Println("Output:", isPowerOfThree(-1))
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
