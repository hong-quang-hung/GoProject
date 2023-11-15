package medium

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
func init() {
	Solutions[1846] = func() {
		fmt.Println("Input: arr = [2,2,1,2,1]")
		fmt.Println("Output:", maximumElementAfterDecrementingAndRearranging([]int{2, 2, 1, 2, 1}))
		fmt.Println("Input: arr = [100,1,1000]")
		fmt.Println("Output:", maximumElementAfterDecrementingAndRearranging([]int{100, 1, 1000}))
		fmt.Println("Input: arr = [1,2,3,4,5]")
		fmt.Println("Output:", maximumElementAfterDecrementingAndRearranging([]int{1, 2, 3, 4, 5}))
	}
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	count := make([]int, n+1)
	for _, a := range arr {
		count[min(a, n)]++
	}

	res := 1
	for i := 2; i <= n; i++ {
		res = min(res+count[i], i)
	}
	return res
}
