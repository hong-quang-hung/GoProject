package easy

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/last-stone-weight/
func Leetcode_Last_Stone_Weight() {
	fmt.Println("Input: stones = [2,7,4,1,8,1]")
	fmt.Println("Output:", lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println("Input: heights = [5,5]")
	fmt.Println("Output:", lastStoneWeight([]int{5, 5}))
}

type lastStoneWeightHeap []int

func (h lastStoneWeightHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h lastStoneWeightHeap) Len() int           { return len(h) }
func (h lastStoneWeightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h lastStoneWeightHeap) Empty() bool        { return len(h) == 0 }
func (h *lastStoneWeightHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *lastStoneWeightHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func lastStoneWeight(stones []int) int {
	stoneHeap := new(lastStoneWeightHeap)
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
