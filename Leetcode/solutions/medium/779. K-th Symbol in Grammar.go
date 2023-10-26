package medium

import (
	"fmt"
	"math/bits"
)

// Reference: https://leetcode.com/problems/k-th-symbol-in-grammar/
func init() {
	Solutions[779] = func() {
		fmt.Println("Input: n = 1, k = 1")
		fmt.Println("Output:", kthGrammar(1, 1))
		fmt.Println("Input: n = 2, k = 2")
		fmt.Println("Output:", kthGrammar(2, 2))
	}
}

func kthGrammar(n int, k int) int {
	return bits.OnesCount(uint(k-1)) % 2
}
