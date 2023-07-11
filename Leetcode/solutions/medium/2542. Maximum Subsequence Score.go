package medium

import (
	"container/heap"
	"fmt"
	"sort"
)

// Reference: https://leetcode.com/problems/maximum-subsequence-score/
func init() {
	Solutions[2542] = func() {
		fmt.Println("Input: nums1 = [1,3,3,2], nums2 = [2,1,3,4], k = 3")
		fmt.Println("Output:", maxScore_ii([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3))
		fmt.Println("Input: nums1 = [4,2,3,1,1], nums2 = [7,5,10,9,6], k = 1")
		fmt.Println("Output:", maxScore_ii([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1))
		fmt.Println("Input: nums1 = [22,5,25,15,28,1], nums2 = [11,7,13,6], k = 3")
		fmt.Println("Output:", maxScore_ii([]int{22, 5, 25, 15, 28, 1}, []int{22, 30, 25, 25, 9, 18}, 3))
	}
}

func maxScore_ii(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	item := make([][2]int, n)
	for i := 0; i < n; i++ {
		item[i] = [2]int{nums1[i], nums2[i]}
	}

	sort.Slice(item, func(i, j int) bool {
		return item[i][1] > item[j][1]
	})

	h, topKSum := new(MinHeap), int64(0)
	for i := 0; i < k; i++ {
		topKSum += int64(item[i][0])
		heap.Push(h, item[i][0])
	}

	res := int64(topKSum) * int64(item[k-1][1])
	for i := k; i < n; i++ {
		topKSum += int64(item[i][0]) - int64(heap.Pop(h).(int))
		res = max64(res, topKSum*int64(item[i][1]))
		heap.Push(h, item[i][0])
	}
	return res
}
