package easy

import "fmt"

func init() {
	Solutions[342] = Leetcode_Is_Power_Of_Four
}

// Reference: https://leetcode.com/problems/power-of-three/
func Leetcode_Is_Power_Of_Four() {
	fmt.Println("Input: n = 16")
	fmt.Println("Output:", isPowerOfFour(16))
	fmt.Println("Input: n = -1")
	fmt.Println("Output:", isPowerOfFour(-1))
}

func isPowerOfFour(n int) bool {
	return (n > 0) && ((n & (n - 1)) == 0) && (n%3 == 1)
}
