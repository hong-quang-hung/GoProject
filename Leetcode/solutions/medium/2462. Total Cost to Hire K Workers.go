package medium

import (
	"container/heap"
	"fmt"
)

func init() {
	Solutions[2462] = Leetcode_Total_Cost
}

// Reference: https://leetcode.com/problems/total-cost-to-hire-k-workers/
func Leetcode_Total_Cost() {
	fmt.Println("Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4")
	fmt.Println("Output:", totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	fmt.Println("Input: costs = [1,2,4,1], k = 3, candidates = 3")
	fmt.Println("Output:", totalCost([]int{1, 2, 4, 1}, 3, 3))
	fmt.Println("Input: costs = [2,2,2,2,2,2,1,4,5,5,5,5,5,2,2,2,2,2,2,2,2,2,2,2,2,2], k = 7, candidates = 3")
	fmt.Println("Output:", totalCost([]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 7, 3))
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := new(MinHeap)
	rightHeap := new(MinHeap)
	left, right := 0, len(costs)-1
	for i := 0; i < candidates && left <= right; i++ {
		heap.Push(leftHeap, costs[left])
		left++
	}

	for i := 0; i < candidates && left <= right; i++ {
		heap.Push(rightHeap, costs[right])
		right--
	}

	res := int64(0)
	for i := 0; i < k; i++ {
		if rightHeap.Len() == 0 || (leftHeap.Len() > 0 && leftHeap.Peek() <= rightHeap.Peek()) {
			res += int64(heap.Pop(leftHeap).(int))
			if left <= right {
				heap.Push(leftHeap, costs[left])
				left++
			}
		} else {
			res += int64(heap.Pop(rightHeap).(int))
			if left <= right {
				heap.Push(rightHeap, costs[right])
				right--
			}
		}
	}
	return res
}
