package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/
func init() {
	Solutions[1611] = func() {
		fmt.Println("Input: n = 5")
		fmt.Println("Output:", minimumOneBitOperations(5))
		fmt.Println("Input: n = 6")
		fmt.Println("Output:", minimumOneBitOperations(6))
	}
}

func minimumOneBitOperations(n int) int {
	res := n
	n = n >> 1
	for n > 0 {
		res ^= n
		n = n >> 1
	}
	return res
}
