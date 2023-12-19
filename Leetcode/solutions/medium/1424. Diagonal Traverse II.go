package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/diagonal-traverse-ii/
func init() {
	Solutions[1424] = func() {
		fmt.Println(`Input: nums = [[1,2,3],[4,5,6],[7,8,9]]`)
		fmt.Println(`Output:`, findDiagonalOrder(S2SoSliceInt(`[[1,2,3],[4,5,6],[7,8,9]]`)))
		fmt.Println(`Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]`)
		fmt.Println(`Output:`, findDiagonalOrder(S2SoSliceInt(`[[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]`)))
	}
}

type DiagonalHeap [][2]int

func (h DiagonalHeap) Less(i, j int) bool {
	sumI := h[i][0] + h[i][1]
	sumJ := h[j][0] + h[j][1]
	return sumI < sumJ || sumI == sumJ && h[i][0] > h[j][0]
}
func (h DiagonalHeap) Len() int      { return len(h) }
func (h DiagonalHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h DiagonalHeap) Empty() bool   { return len(h) == 0 }
func (h *DiagonalHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *DiagonalHeap) Push(i interface{}) {
	*h = append(*h, i.([2]int))
}

func findDiagonalOrder(nums [][]int) []int {
	h := new(DiagonalHeap)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			heap.Push(h, [2]int{i, j})
		}
	}

	res := []int{}
	for h.Len() > 0 {
		item := heap.Pop(h).([2]int)
		res = append(res, nums[item[0]][item[1]])
	}
	return res
}
