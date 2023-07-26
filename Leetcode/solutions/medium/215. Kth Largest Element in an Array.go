package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/single-number/
func init() {
	Solutions[215] = func() {
		fmt.Println("Input: nums = [3,2,1,5,6,4], k = 2")
		fmt.Println("Output:", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
		fmt.Println("Input: nums = [3,2,3,1,2,4,5,5,6], k = 4")
		fmt.Println("Output:", findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	}
}

func findKthLargest(nums []int, k int) int {
	h := new(MaxHeap)
	for _, num := range nums {
		heap.Push(h, num)
	}

	var res int
	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}
