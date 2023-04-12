package hard

import (
	"container/heap"
	"fmt"
	"math"

	"leetcode.com/Leetcode/utils"
)

// Reference: https://leetcode.com/problems/minimize-deviation-in-array/
func Leetcode_Minimum_Deviation() {
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println("Input: nums = [4,1,5,20,3]")
	fmt.Println("Output:", minimumDeviation(utils.Slice(4, 1, 5, 20, 3)))
}

type minimumDeviationHeap []int

func (h minimumDeviationHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h minimumDeviationHeap) Len() int           { return len(h) }
func (h minimumDeviationHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h minimumDeviationHeap) Empty() bool        { return len(h) == 0 }
func (h *minimumDeviationHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *minimumDeviationHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}

func minimumDeviation(nums []int) int {
	set := new(minimumDeviationHeap)
	minDev := math.MaxInt32
	for _, num := range nums {
		if num%2 == 1 {
			num *= 2
		}

		minDev = min(minDev, num)
		heap.Push(set, num)
	}

	res := math.MaxInt32
	for !set.Empty() {
		pop := heap.Pop(set).(int)
		res = min(res, pop-minDev)
		if pop%2 == 0 {
			pop /= 2
			minDev = min(minDev, pop)
			heap.Push(set, pop)
		} else {
			break
		}
	}
	return res
}
