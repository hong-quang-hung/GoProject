package hard

import (
	"fmt"
)

// Reference: https://leetcode.com/problems/sliding-window-maximum/
func init() {
	Solutions[239] = func() {
		fmt.Println("Input: nums = [1,3,-1,-3,5,3,6,7], k = 3")
		fmt.Println("Output:", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
		fmt.Println("Input: nums = [1], k = 1")
		fmt.Println("Output:", maxSlidingWindow([]int{1}, 1))
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}
	res := []int{}
	for i, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i < k-1 {
			continue
		}

		res = append(res, nums[deque[0]])
		if deque[0] == i-k+1 {
			deque = deque[1:]
		}
	}
	return res
}
