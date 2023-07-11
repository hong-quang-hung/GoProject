package medium

import (
	"fmt"
	"math/bits"
)

// Reference: https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/
func init() {
	Solutions[1318] = func() {
		fmt.Println("Input: a = 2, b = 6, c = 5")
		fmt.Println("Output:", minFlips(2, 6, 5))
		fmt.Println("Input: a = 4, b = 2, c = 7")
		fmt.Println("Output:", minFlips(4, 2, 7))
		fmt.Println("Input: a = 1, b = 2, c = 3")
		fmt.Println("Output:", minFlips(1, 2, 3))
	}
}

func minFlips(a int, b int, c int) int {
	n := max(bits.Len(uint(a)), max(bits.Len(uint(b)), bits.Len(uint(c))))
	res := 0
	for i := 0; i < n; i++ {
		bitA, bitB, bitC := a%2, b%2, c%2
		if (bitA | bitB) != bitC {
			if bitA == 1 && bitB == 1 && bitC == 0 {
				res += 2
			} else {
				res += 1
			}
		}

		a = a >> 1
		b = b >> 1
		c = c >> 1
	}
	return res
}
