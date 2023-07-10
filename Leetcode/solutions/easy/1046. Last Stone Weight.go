package easy

import (
	"container/heap"
	"fmt"
)

func init() {
	Solutions[1046] = Leetcode_Last_Stone_Weight
}

// Reference: https://leetcode.com/problems/last-stone-weight/
func Leetcode_Last_Stone_Weight() {
	fmt.Println("Input: stones = [2,7,4,1,8,1]")
	fmt.Println("Output:", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("Input: heights = [5,5]")
	fmt.Println("Output:", lastStoneWeight([]int{5, 5}))
}

func lastStoneWeight(stones []int) int {
	stoneHeap := new(MaxHeap)
	for _, stone := range stones {
		*stoneHeap = append(*stoneHeap, stone)
	}

	heap.Init(stoneHeap)
	for stoneHeap.Len() > 1 {
		y := heap.Pop(stoneHeap).(int)
		x := heap.Pop(stoneHeap).(int)
		z := y - x
		if z > 0 {
			heap.Push(stoneHeap, z)
		}
	}

	if stoneHeap.Len() == 0 {
		return 0
	}
	return heap.Pop(stoneHeap).(int)
}
