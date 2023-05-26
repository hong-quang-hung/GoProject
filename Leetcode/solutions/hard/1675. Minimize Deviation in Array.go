package hard

import (
	"container/heap"
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimize-deviation-in-array/
func Leetcode_Minimum_Deviation() {
	fmt.Println("Input: nums = [1,2,3,4]")
	fmt.Println("Output:", minimumDeviation([]int{1, 2, 3, 4}))
	fmt.Println("Input: nums = [4,1,5,20,3]")
	fmt.Println("Output:", minimumDeviation([]int{4, 1, 5, 20, 3}))
}

func minimumDeviation(nums []int) int {
	set := new(MaxHeap)
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
