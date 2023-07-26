package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-speed-to-arrive-on-time/
func init() {
	Solutions[1870] = func() {
		fmt.Println("Input: dist = [1,3,2], hour = 6")
		fmt.Println("Output:", minSpeedOnTime([]int{1, 3, 2}, 6))
		fmt.Println("Input: dist = [1,3,2], hour = 2.7")
		fmt.Println("Output:", minSpeedOnTime([]int{1, 3, 2}, 2.7))
		fmt.Println("Input: dist = [1,3,2], hour = 1.9")
		fmt.Println("Output:", minSpeedOnTime([]int{1, 3, 2}, 1.9))
	}
}

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	res := -1
	left, right := 1, 10_000_000
	for left <= right {
		mid := (left + right) / 2
		time := float64(0)
		for i := 0; i < n-1; i++ {
			time += math.Ceil(float64(dist[i]) / float64(mid))
		}
		time += float64(dist[n-1]) / float64(mid)

		if time <= hour {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
