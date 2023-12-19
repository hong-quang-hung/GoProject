package easy

import "fmt"

// Reference: https://leetcode.com/problems/counting-bits/
func init() {
	Solutions[338] = func() {
		fmt.Println(`Input: n = 2`)
		fmt.Println(`Output:`, countBits(2))
		fmt.Println(`Input: n = 5`)
		fmt.Println(`Output:`, countBits(5))
	}
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1) // bits.OnesCount(uint(i))
	}
	return res
}
