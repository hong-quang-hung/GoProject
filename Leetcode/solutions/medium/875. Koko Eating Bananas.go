package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/koko-eating-bananas/
func init() {
	Solutions[875] = func() {
		fmt.Println("Input: piles = [3,6,7,11], h = 8")
		fmt.Println("Output:", minEatingSpeed([]int{3, 6, 7, 11}, 8))
		fmt.Println("Input: ppiles = [30,11,23,4,20], h = 6")
		fmt.Println("Output:", minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
	}
}

func minEatingSpeed(piles []int, h int) int {
	var left int = 1
	var right int = 0
	for _, p := range piles {
		if right < p {
			right = p
		}
	}
	var res int = right
	for left <= right {
		var hours int64 = 0
		var mid = left + (right-left)/2
		for _, p := range piles {
			hours += int64(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= int64(h) {
			if res > mid {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
