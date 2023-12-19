package easy

import "fmt"

// Reference: https://leetcode.com/problems/power-of-three/
func init() {
	Solutions[326] = func() {
		fmt.Println(`Input: n = 27`)
		fmt.Println(`Output:`, isPowerOfThree(27))
		fmt.Println(`Input: n = -1`)
		fmt.Println(`Output:`, isPowerOfThree(-1))
	}
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
