package medium

import (
	"container/heap"
	"fmt"
)

// Reference: https://leetcode.com/problems/top-k-frequent-elements/
func init() {
	Solutions[347] = func() {
		fmt.Println(`Input: nums = [1,1,1,2,2,3], k = 2`)
		fmt.Println(`Output:`, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
		fmt.Println(`Input: nums = [1], k = 1`)
		fmt.Println(`Output:`, topKFrequent([]int{1}, 1))
	}
}

type topKFrequentHeap [][2]int

func (h topKFrequentHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h topKFrequentHeap) Len() int           { return len(h) }
func (h topKFrequentHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h topKFrequentHeap) Empty() bool        { return len(h) == 0 }
func (h *topKFrequentHeap) Pop() interface{} {
	r := (*h)[(*h).Len()-1]
	*h = (*h)[0 : (*h).Len()-1]
	return r
}
func (h *topKFrequentHeap) Push(i interface{}) {
	*h = append(*h, i.([2]int))
}

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	h := new(topKFrequentHeap)
	for k, v := range m {
		heap.Push(h, [2]int{k, v})
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		val := heap.Pop(h).([2]int)
		res[i] = val[0]
	}
	return res
}
