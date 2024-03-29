package medium

import (
	"fmt"
	"math"
)

// Reference: https://leetcode.com/problems/minimum-time-to-complete-trips/
func init() {
	Solutions[2187] = func() {
		fmt.Println(`Input: time = [1,2,3], totalTrips = 5`)
		fmt.Println(`Output:`, minimumTime([]int{1, 2, 3}, 5))
	}
}

func minimumTime(time []int, totalTrips int) int64 {
	left, right := int64(0), int64(math.MaxInt64)
	for left < right {
		mid := left + (right-left)/2
		trips := int64(0)
		for _, t := range time {
			trips += mid / int64(t)
			if trips >= int64(totalTrips) {
				break
			}
		}

		if trips < int64(totalTrips) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
