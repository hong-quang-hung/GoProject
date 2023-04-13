package medium

import "fmt"

// Reference: https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/
func Leetcode_Ship_Within_Days() {
	fmt.Println("Input: weights = [1,2,3,4,5,6,7,8,9,10], days = 5")
	fmt.Println("Output:", shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println("Input: weights = [3,2,2,4,1,4], days = 3")
	fmt.Println("Output:", shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
	fmt.Println("Input: weights = [1,2,3,1,1], days = 4")
	fmt.Println("Output:", shipWithinDays([]int{1, 2, 3, 1, 1}, 4))
}

func shipWithinDays(weights []int, days int) int {
	maxWeight, total := 0, 0
	for _, weight := range weights {
		total += weight
		maxWeight = max(maxWeight, weight)
	}

	left, right := maxWeight, total
	for left < right {
		mid := (left + right) / 2
		if feasible(weights, mid, days) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func feasible(weights []int, mid, days int) bool {
	daysNeeded, currentLoad := 1, 0
	for _, weight := range weights {
		currentLoad += weight
		if currentLoad > mid {
			daysNeeded++
			currentLoad = weight
		}
	}
	return daysNeeded <= days
}
