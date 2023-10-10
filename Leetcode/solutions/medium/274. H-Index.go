package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/h-index/
func init() {
	Solutions[274] = func() {
		fmt.Println("Input: citations = [3,0,6,1,5]")
		fmt.Println("Output:", hIndex([]int{3, 0, 6, 1, 5}))
		fmt.Println("Input: citations = [1,3,1]")
		fmt.Println("Output:", hIndex([]int{1, 3, 1}))
	}
}

func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, citation := range citations {
		if citation >= n {
			buckets[n]++
		} else {
			buckets[citation]++
		}
	}

	counter := 0
	for i := n; i >= 0; i-- {
		counter += buckets[i]
		if counter >= i {
			return i
		}
	}
	return 0
}
