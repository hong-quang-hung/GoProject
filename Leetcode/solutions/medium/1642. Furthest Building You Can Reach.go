package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/furthest-building-you-can-reach/
func init() {
	Solutions[1642] = func() {
		fmt.Println(`Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1`)
		fmt.Println(`Output:`, furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
		fmt.Println(`Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2`)
		fmt.Println(`Output:`, furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
		fmt.Println(`Input: heights = [14,3,19,3], bricks = 17, ladders = 0`)
		fmt.Println(`Output:`, furthestBuilding([]int{14, 3, 19, 3}, 17, 0))
	}
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	h := new(MinHeap)
	for i := 1; i < n; i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}

		heap.Push(h, diff)
		if h.Len() > ladders {
			bricks -= heap.Pop(h).(int)
		}

		if bricks < 0 {
			return i - 1
		}
	}
	return n - 1
}
