package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
func init() {
	Solutions[373] = func() {
		fmt.Println(`Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3`)
		fmt.Println(`Output:`, kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
		fmt.Println(`Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2`)
		fmt.Println(`Output:`, kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
		fmt.Println(`Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 10`)
		fmt.Println(`Output:`, kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10))
	}
}

type kSmallestPairsHeap [][]int

func (h kSmallestPairsHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h kSmallestPairsHeap) Len() int           { return len(h) }
func (h kSmallestPairsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h kSmallestPairsHeap) Empty() bool        { return len(h) == 0 }
func (h *kSmallestPairsHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *kSmallestPairsHeap) Push(i interface{}) {
	*h = append(*h, i.([]int))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := new(kSmallestPairsHeap)
	heap.Push(h, []int{nums1[0] + nums2[0], 0, 0})
	res := [][]int{}
	for h.Len() != 0 && len(res) < k {
		cur := heap.Pop(h).([]int)
		i, j := cur[1], cur[2]
		res = append(res, []int{nums1[i], nums2[j]})

		if j == 0 && i+1 < len(nums1) {
			heap.Push(h, []int{nums1[i+1] + nums2[0], i + 1, 0})
		}
		if j+1 < len(nums2) {
			heap.Push(h, []int{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return res
}
